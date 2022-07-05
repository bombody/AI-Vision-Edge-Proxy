package mqtt

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/chryscloud/video-edge-ai-proxy/utils"
)

// Removes a camera from the edge
func (mqtt *mqttManager) StopCamera(configPayload []byte) error {
	g.Log.Info("received