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
func (gih *grpcImageHandler) Annotate(ctx context.Context, req *pb.An