package mqtt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/chryscloud/video-edge-ai-proxy/utils"
	badger "github.com/dgraph-io/badger/v2"
	"github.com/docker/docker/api/types/events"
)

const (
	ProcessActionDie   = "die"
	ProcessActionStart = "start"
)

type ProcessState struct {
	Time     int64  // seconds since epoch
	DeviceID string // deviceID
	Action   string // process action from docker events
}

// Check settings and also if MQTT initial connection has been made
func (mqtt *mqttManager) getMQTTSettings() (*models.Settings, error) {
	// check settings if they exist
	settings, err := mqtt.settingsService.Get()
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, ErrNoMQTTSettings
		}
		g.Log.Error("failed to retrieve edge settings", err)
		return nil, err
	}
	if settings.ProjectID == "" || settings.Region == "" || settings.GatewayID == "" || settings.RegistryID == "" || settings.PrivateRSAKey == nil {
		return nil, ErrNoMQTTSettings
	}
	return settings, nil
}

// config and commans subscribers
func (mqtt *mqttManager) gatewaySubscribers() error {
	// wait for connection to be opened and propagate

	errBind := mqtt.bindAllDevices()
	if errBind != nil {
		g.Log.Error("failed to report bind devices", errBind)
		return errBind
	}

	errCfg := mqtt.subscribeToConfig(mqtt.