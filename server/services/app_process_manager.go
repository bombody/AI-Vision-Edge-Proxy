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
	"strconv"
	"strings"
	"time"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	dockerErrors "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/go-redis/redis/v7"
)

// ProcessManager - start, stop of docker containers
type AppProcessManager struct {
	storage *Storage
	rdb     *redis.Client
}

func NewAppManager(storage *Storage, rdb *redis.Client) *AppProcessManager {
	return &AppProcessManager{
		storage: storage,
		rdb:     rdb,
	}
}

// Install - installs the new app
func (am *AppProcessManager) Install(app *models.AppProcess) (*models.AppProcess, error) {

	// installation process
	cl := docker.NewSocketClient(docker.Log(g.Log), docker.Host("unix:///var/run/docker.sock"))

	fl := filters.NewArgs()
	pruneReport, pruneErr := cl.ContainersPrune(fl)
	if pruneErr != nil {
		g.Log.Error("container prunning fialed", pruneErr)
		return nil, pruneErr
	}
	g.Log.Info("app prune successfull. Report and space reclaimed", pruneReport.ContainersDeleted, pruneReport.SpaceReclaimed)

	// expose desired ports mappings if any
	portMap := nat.PortMap{}
	portSet := nat.PortSet{}
	if len(app.PortMapping) > 0 {

		for _, pm := range app.PortMapping {
			exposedPort := strconv.Itoa(pm.Exposed)
			mapsTo := pm.MapTo

			mapsToPort := strconv.Itoa(mapsTo) + "/tcp"
			portSet[nat.Port(mapsToPort)] = struct{}{}
			portMap[nat.Port(mapsToPort)] = []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: exposedPort}}
		}
	}

	// prepare host configuration
	hostConfig := &container.HostConfig{
		LogConfig: container.LogConfig