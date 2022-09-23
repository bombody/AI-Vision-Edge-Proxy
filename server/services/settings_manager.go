
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

package services

import (
	"encoding/json"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/chryscloud/go-microkit-plugins/docker"
	dockerhub "github.com/chryscloud/go-microkit-plugins/dockerhub"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/chryscloud/video-edge-ai-proxy/utils"
	"github.com/dgraph-io/badger/v2"
	"github.com/docker/docker/api/types"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-version"
)

// SettingsManager - various settings for the edge
type SettingsManager struct {
	storage             *Storage
	current_edge_key    string
	current_edge_secret string
	mux                 *sync.RWMutex
	apiClient           *resty.Client
}

func NewSettingsManager(storage *Storage) *SettingsManager {
	return &SettingsManager{
		storage:   storage,
		mux:       &sync.RWMutex{},
		apiClient: resty.New(),
	}
}

func (sm *SettingsManager) GetCurrentEdgeKeyAndSecret() (string, string, error) {
	if sm.current_edge_key == "" || sm.current_edge_secret == "" {
		settings, err := sm.getDefault()
		if err != nil {
			if err != badger.ErrKeyNotFound {
				g.Log.Error("failed to query for current edge api key and secret", err)
			}
			return "", "", err
		}
		sm.mux.Lock()
		defer sm.mux.Unlock()
		sm.current_edge_key = settings.EdgeKey
		sm.current_edge_secret = settings.EdgeSecret
	}
	return sm.current_edge_key, sm.current_edge_secret, nil
}

// Used on systm start, calling cloud to connect to (and refresh possible keys, cert, ...)
func (sm *SettingsManager) UpdateEdgeRegistrationToCloud() error {
	defaultSettings, err := sm.getDefault()
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil
		}
		g.Log.Error("failed to query for settings in datastore", err)
		return err
	}

	// if not connected to edge there is nothing to do
	if defaultSettings.EdgeKey == "" || defaultSettings.EdgeSecret == "" {
		return nil
	}

	// get system info
	sysInfo, err := sm.GetSystemInfo()
	if err != nil {
		g.Log.Error("failed to retrieve system info", err)
		return err
	}

	_, err = sm.updateSettingsWithMQTTCredentials(sysInfo, defaultSettings)
	return err
}

// getDefault - retrieves settings if exist, otherwise creates new empty settings
func (sm *SettingsManager) getDefault() (*models.Settings, error) {
	// check if settings already exist
	settingsBytes, err := sm.storage.Get(models.PrefixSettingsKey, models.SettingDefaultKey)
	if err != nil {
		if err != badger.ErrKeyNotFound {
			g.Log.Error("failed to retrieve settings", err)
			return nil, err
		}
	}

	var settings models.Settings
	if settingsBytes == nil {
		settings = models.Settings{
			Name: models.SettingDefaultKey,
		}
	} else {
		unmErr := json.Unmarshal(settingsBytes, &settings)
		if unmErr != nil {
			g.Log.Error("failed to unmarshal settings", unmErr)
			return nil, unmErr
		}
	}
	sm.mux.Lock()
	defer sm.mux.Unlock()
	if settings.EdgeKey != "" {
		sm.current_edge_key = settings.EdgeKey
	}
	if settings.EdgeSecret != "" {
		sm.current_edge_secret = settings.EdgeSecret
	}
	return &settings, nil
}

// Overwrite always overwrites the complete settings
func (sm *SettingsManager) Overwrite(settings *models.Settings) (*models.Settings, error) {
	existingSettings, _ := sm.getDefault()

	// get system info
	sysInfo, err := sm.GetSystemInfo()
	if err != nil {
		g.Log.Error("failed to retrieve system info", err)
		return nil, err
	}

	if existingSettings != nil {
		sysInfo.GatewayID = existingSettings.GatewayID
		sysInfo.RegistryID = existingSettings.RegistryID
	}

	updSettings, err := sm.updateSettingsWithMQTTCredentials(sysInfo, settings)
	if err != nil {
		g.Log.Error("failed to update settings", err)
		return nil, err
	}

	return updSettings, nil
}

func (sm *SettingsManager) updateSettingsWithMQTTCredentials(sysInfo *models.SystemInfo, settings *models.Settings) (*models.Settings, error) {
	// validate settings with the Chrysalis Cloud

	if settings.GatewayID != "" && settings.RegistryID != "" {
		sysInfo.GatewayID = settings.GatewayID
		sysInfo.RegistryID = settings.RegistryID
	}
	var resp []byte
	for i := 0; i < 3; i++ {
		response, apiErr := utils.CallAPIWithBody(sm.apiClient, "POST", g.Conf.API.Endpoint+"/api/v1/edge/credentials", sysInfo, settings.EdgeKey, settings.EdgeSecret)
		if apiErr != nil {
			if apiErr == models.ErrForbidden {
				g.Log.Warn("authentication failed communicating with chrysalis cloud", apiErr)
				return nil, apiErr
			}
			if apiErr == models.ErrProcessNotFound {
				// if not found, then try without gatewayID and registryID
				sysInfo.GatewayID = ""
				sysInfo.RegistryID = ""
				continue
			}
			g.Log.Error("Failed to validate credentials with chrys cloud", apiErr)
			// AbortWithError(c, http.StatusUnauthorized, "Failed to validate credentials with Chryscloud")
			return nil, apiErr
		}
		resp = response
		break
	}

	var cloudResponse models.EdgeConnectCredentials
	mErr := json.Unmarshal(resp, &cloudResponse)
	if mErr != nil {
		g.Log.Error("failed to unmarshal response from Chryscloud", mErr)
		// AbortWithError(c, http.StatusExpectationFailed, "Failed to unmarshal response from Chryscloud. Please upgrade Chrysalis Edge Proxy to latest version")
		return nil, mErr
	}
	settings.ProjectID = cloudResponse.ProjectID
	settings.RegistryID = cloudResponse.RegistryID
	settings.GatewayID = cloudResponse.GatewayID
	settings.Region = cloudResponse.Region
	settings.PrivateRSAKey = cloudResponse.PrivateKeyPem