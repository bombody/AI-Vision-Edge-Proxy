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
	g.Log.Info("received payload to remove a device")
	var payload models.EdgeCommandPayload
	err := json.Unmarshal(configPayload, &payload)
	if err != nil {
		g.Log.Error("failed to unmarshal config payload", err)
		return err
	}
	_, pErr := mqtt.processService.Info(payload.Name)
	if pErr != nil {
		if pErr == models.ErrProcessNotFound {
			// nothing to do, but report unbinding
			e