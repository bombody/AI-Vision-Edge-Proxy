package grpcapi

import (
	"context"
	"time"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	pb "github.com/chryscloud/video-edge-ai-proxy/proto"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Annotate queues a new annotation event to be sent to Chrysalis event servers
func (gih *grpcImageHandler) Annotate(ctx context.Context, req *pb.AnnotateRequest) (*pb.AnnotateResponse, error) {
	if gih.edgeKey == nil {
		settings, err := gih.settingsManager.Get()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to read settings")
		}
		if settings.EdgeKey == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Can't find edge key in settings. required to use annotations. Visit https://cloud.chryscloud.com to enable annotations and storage capabilities from the edge.")
		}
		gih.edgeKey = &settings.EdgeKey
	}
	weekPast := time.Now().AddDate(0, 0, -7).Unix() * 1000
	weekFuture := time.Now().AddDate(0, 0, 7).Unix() * 1000
	if req.DeviceName == "" || req.Type == "" || req.StartTimestamp < 0 {
		return nil, status.Errorf(codes.Invalid