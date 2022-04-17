package api

import (
	"net/http"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/chryscloud/video-edge-ai-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

type testApiHandler struct {
	rdb *redis.Client
}

func NewTestApiHandler(rdb *redis.Client) *testApiHan