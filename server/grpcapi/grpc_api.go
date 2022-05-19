
// Copyright 2020 Wearless Tech Inc All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcapi

import (
	"container/list"
	"context"
	"encoding/base64"
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adjust/rmq/v2"
	"github.com/chryscloud/video-edge-ai-proxy/batch"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	pb "github.com/chryscloud/video-edge-ai-proxy/proto"
	"github.com/chryscloud/video-edge-ai-proxy/services"
	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/proto"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageInput struct {
	Enable bool `json:"enable"`
}

type grpcImageHandler struct {
	redisConn               *redis.Client
	deviceMap               sync.Map
	processManager          *services.ProcessManager
	settingsManager         *services.SettingsManager
	edgeKey                 *string
	msgQueue                rmq.Queue
	realtimeCache           sync.Map
	realtimeDeviceQueryTime sync.Map
}

// NewGrpcImageHandler returns main GRPC API handler
func NewGrpcImageHandler(processManager *services.ProcessManager, settingsManager *services.SettingsManager, rdb *redis.Client) *grpcImageHandler {

	conn := rmq.OpenConnectionWithRedisClient("annotationService", rdb)
	msgQueue := conn.OpenQueue("annotationqueue")

	// add batch listener (consumer) for annotatons
	annotationConsumer := batch.NewAnnotationConsumer(0, settingsManager, msgQueue)
	msgQueue.StartConsuming(g.Conf.Annotation.UnackedLimit, time.Duration(g.Conf.Annotation.PollDurationMs)*time.Millisecond)
	msgQueue.AddBatchConsumerWithTimeout("annotationqueue", g.Conf.Annotation.MaxBatchSize, time.Duration(g.Conf.Annotation.PollDurationMs)*time.Millisecond, annotationConsumer)

	return &grpcImageHandler{
		redisConn:               rdb,
		deviceMap:               sync.Map{},
		processManager:          processManager,
		settingsManager:         settingsManager,
		msgQueue:                msgQueue,
		realtimeCache:           sync.Map{},
		realtimeDeviceQueryTime: sync.Map{},
	}
}

func (gih *grpcImageHandler) toUint64(object map[string]interface{}, field string) int64 {
	if val, ok := object[field]; ok {
		strVal := val.(string)
		w, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			g.Log.Error("Failed to convert width to int", err)
		}
		return w
	}
	return 0
}

// ListStreams returns the list of all streams regardless of their status
func (gih *grpcImageHandler) ListStreams(req *pb.ListStreamRequest, stream pb.Image_ListStreamsServer) error {
	err := gih.processManager.ListStream(stream.Context(), func(process *models.StreamProcess) error {
		res := &pb.ListStream{
			Name:       process.Name,
			Dead:       process.State.Dead,
			Error:      process.State.Error,
			ExitCode:   int64(process.State.ExitCode),
			Oomkilled:  process.State.OOMKilled,
			Paused:     process.State.Paused,
			Pid:        int32(process.State.Pid),
			Restarting: process.State.Restarting,
			Running:    process.State.Running,
			Status:     process.Status,
		}
		if process.State.Health != nil {
			res.FailingStreak = int64(process.State.Health.FailingStreak)
			res.HealthStatus = process.State.Health.Status
		}

		err := stream.Send(res)
		if err != nil {
			g.Log.Error("failed to send process item", err)
			return err
		}
		g.Log.Info("sent process with name: ", process.Name)
		return nil
	})
	if err != nil {
		g.Log.Error("failed to retrieve processes stream", err)
	}
	return nil
}

// VideoLatestImage - bidirectional connection with client continously sending live video image
func (gih *grpcImageHandler) VideoLatestImage(ctx context.Context, request *pb.VideoFrameRequest) (*pb.VideoFrame, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	streamName := request.DeviceId

	// every 5 seconds report last query time
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	lastQTime := int64(0)
	if lastDeviceQueryTime, ok := gih.realtimeDeviceQueryTime.Load(request.DeviceId); ok {
		lastQTime = lastDeviceQueryTime.(int64)
	}
	if currentTime-lastQTime > (5000) {
		// g.Log.Info("updating last query time for", request.DeviceId)
		// storeValuesStart := time.Now().UnixNano()

		isKeyFrameOnly := request.KeyFrameOnly

		decodeOnlyKeyFramesKey := models.RedisIsKeyFrameOnlyPrefix + streamName
		err := gih.redisConn.Set(decodeOnlyKeyFramesKey, strconv.FormatBool(isKeyFrameOnly), 0).Err()
		if err != nil {
			g.Log.Error("failed to set if is keyframe only", streamName, err)
			return nil, status.Errorf(codes.Internal, "failed to set preferences in redis")
		}

		valMap := make(map[string]interface{}, 0)
		valMap[models.RedisLastAccessQueryTimeKey] = currentTime

		rErr := gih.redisConn.HSet(models.RedisLastAccessPrefix+streamName, valMap).Err()
		if rErr != nil {
			g.Log.Error("failed to update on stopProxy redis", streamName, rErr)
			return nil, status.Errorf(codes.Internal, "can't access redis")
		}

		// storeValueEnd := time.Now().UnixNano()
		// g.Log.Info("Time to store query values [ms] ", (storeValueEnd-storeValuesStart)/1e+6)
		gih.realtimeDeviceQueryTime.Store(request.DeviceId, currentTime)
	}

	// // loading VideoFrame from redis
	vf := &pb.VideoFrame{}

	isDeviceFirstRun := false

	for i := 0; i < 3; i++ {

		var cache *list.List
		if cacheVal, ok := gih.realtimeCache.Load(request.DeviceId); ok {
			cache = cacheVal.(*list.List)
		} else {
			cache = list.New()
			gih.realtimeCache.Store(request.DeviceId, cache)
			isDeviceFirstRun = true
		}
		if isDeviceFirstRun {
			go func() {
				gih.cacheLiveVideo(request.DeviceId)
			}()

			time.Sleep(time.Millisecond * 200)
		}

		if cache.Len() > 0 {
			// g.Log.Info("Cache length: ", request.DeviceId, cache.Len())
			front := cache.Front()
			if front != nil {
				redisVal := front.Value.(redis.XMessage)
				vf = gih.unmarshalRedisImage(vf, request.DeviceId, redisVal)

				// if cache.Len() > 1 {
				cache.Remove(front)
				// }
				break