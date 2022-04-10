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

package api

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/chryscloud/video-edge-ai-proxy/services"
	"github.com/chryscloud/video-edge-ai-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v7"
)

type rtspProcessHandler struct {
	processManager  *services.ProcessManager
	settingsManager *services.SettingsManager
	rdb             *redis.Client
}

func NewRTSPProcessHandler(rdb *redis.Client, processManager *services.ProcessManager, settingsManager *services.SettingsManager) *rtspProcessHandler {
	return &rtspProcessHandler{
		processManager:  processManager,
		settingsManager: settingsManager,
		rdb:             rdb,
	}
}

func (ph *rtspProcessHandler) StartRTSP(c *gin.Context) {
	var streamProcess models.StreamProcess
	if err := c.ShouldBindWith(&streamProcess, binding.JSON); err != nil {
		g.Log.Warn("missing required fields", err)
		AbortWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if streamProcess.RTSPEndpoint == "" {
		AbortWithError(c, http.StatusBadRequest, "RTP endpoint required")
		return
	}
	deviceID := streamProcess.Name
	if streamProcess.Name == "" {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(streamProcess.RTSPEndpoint)))
		deviceID = hash
	}
	streamProcess.RTMPStreamStatus = &models.RTMPStreamStatus{
		Storing:   false,
		Streaming: true,
	}

	rtspImageTag := models.CameraTypeToImageTag["rtsp"]
	currentImagesList, err := ph.settingsManager.ListDockerImages(rtspImageTag)
	if err != nil {
		g.Log.Error("failed to list currently available images", err)
		AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = ph.processManager.Start(&streamProcess, currentImagesList)
	if err != nil {
		g.Log.Warn("failed to start process ", deviceID, err)
		AbortWithError(c, http.StatusConflict, err.Error())
		return
	}
	// publish to chrysalis cloud the change
	utils.PublishToRedis(ph.rdb, deviceID, models.MQTTProcessOperat