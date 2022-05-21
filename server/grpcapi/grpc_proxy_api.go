
package grpcapi

import (
	"context"
	"time"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	pb "github.com/chryscloud/video-edge-ai-proxy/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gih *grpcImageHandler) Proxy(ctx context.Context, req *pb.ProxyRequest) (*pb.ProxyResponse, error) {
	deviceID := req.DeviceId

	if deviceID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "device id required")
	}

	info, err := gih.processManager.Info(deviceID)