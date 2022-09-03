
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
	"context"
	"encoding/json"
	"errors"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/dgraph-io/badger/v2"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	dockerErrors "github.com/docker/docker/client"
	"github.com/go-redis/redis/v7"
)

const (
	// Resource: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
	ArchitectureAmd64 = "amd64"
	ArchitectureArm64 = "arm64"
)

var (
	ArchitectureSuffixMap = map[string]string{ArchitectureAmd64: "", ArchitectureArm64: "-arm64v8"}
)

// ProcessManager - start, stop of docker containers
type ProcessManager struct {
	storage *Storage
	rdb     *redis.Client
}

func NewProcessManager(storage *Storage, rdb *redis.Client) *ProcessManager {
	return &ProcessManager{
		storage: storage,
		rdb:     rdb,
	}
}

// Start - starts the docker container with rtsp, device_id and possibly rtmp environment variables.
// Restarts always when something goes wrong within the streaming process
func (pm *ProcessManager) Start(process *models.StreamProcess, imageUpgrade *models.ImageUpgrade) error {

	// detect architecture
	arch := runtime.GOARCH

	if _, ok := ArchitectureSuffixMap[arch]; !ok {
		return errors.New("architecture currently not supported")
	}

	if process.Name == "" || process.RTSPEndpoint == "" {
		return errors.New("required parameters missing")
	}

	if !imageUpgrade.HasImage && !imageUpgrade.HasUpgrade {
		return errors.New("no camera container found. Please refer to documentation on how to pull a docker image manually")
	}

	settingsTagBytes, err := pm.storage.Get(models.PrefixSettingsDockerTagVersions, "rtsp")
	if err != nil {
		if err == badger.ErrKeyNotFound {

			// if no docker tag version stored in database but image does exist on disk, then store settings docker tag version with that image
			tag := models.CameraTypeToImageTag["rtsp"]
			if imageUpgrade == nil {
				return errors.New("Image not found. Please check the docs and pull the docker image manually.")
			}
			maximumExistingTag := tag + ":" + imageUpgrade.CurrentVersion
			// store to database
			g.Log.Info("maximum existing tag od disk found: ", maximumExistingTag)

			settingsTagVersion := &models.SettingDockerTagVersion{
				CameraType: "rtsp",
				Tag:        tag,
				Version:    imageUpgrade.CurrentVersion,
			}
			stb, sErr := pm.storeSettingsTagVersion(settingsTagVersion)
			if sErr != nil {
				g.Log.Error("failed to store new settings tag version ", sErr)
				return sErr
			}

			settingsTagBytes = stb
		} else {
			g.Log.Error("failed to read rtsp tag from settings", err)
			return err
		}
	}

	var settingsTag models.SettingDockerTagVersion
	err = json.Unmarshal(settingsTagBytes, &settingsTag)
	if err != nil {
		g.Log.Error("failed to unamrshal settings tag", err)
		return err
	}
	process.ImageTag = settingsTag.Tag + ":" + settingsTag.Version

	// Check the latest version that exists on the disk (and if is the same as the one in settings)
	// if is not, correct the latest version stored (most likely user chose to manually deleted the newer version)
	if imageUpgrade.CurrentVersion != settingsTag.Version {
		settingsTag.Version = imageUpgrade.CurrentVersion

		process.ImageTag = imageUpgrade.Name + ":" + imageUpgrade.CurrentVersion

		_, sErr := pm.storeSettingsTagVersion(&settingsTag)
		if sErr != nil {
			g.Log.Error("failed to store new settings tag", sErr, ", image version: ", settingsTag.Tag, settingsTag.Version)
			return sErr
		}
	}

	cl := docker.NewSocketClient(docker.Log(g.Log), docker.Host("unix:///var/run/docker.sock"))

	fl := filters.NewArgs()
	pruneReport, pruneErr := cl.ContainersPrune(fl)
	if pruneErr != nil {
		g.Log.Error("container prunning fialed", pruneErr)
		return pruneErr
	}
	g.Log.Info("prune successfull. Report and space reclaimed:", pruneReport.ContainersDeleted, pruneReport.SpaceReclaimed)

	hostConfig := &container.HostConfig{