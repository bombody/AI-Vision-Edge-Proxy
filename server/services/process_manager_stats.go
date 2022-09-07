package services

import (
	"strings"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
)

// StatsAllProcesses created a statistics object or all running containers (related to edge)
func (pm *ProcessManager) StatsAllProcesses(sett *models.Settings) (*models.AllStreamProcessStats, error) {
	cl := docker.NewSocketClient(docker.Log(g.Log), docker.Host("unix:///var/run/docker.sock"))

	systemInfo, diskUsage, err := cl.SystemWideInfo()

	stats := &models.AllStreamProcessStats{}
	// calculate disk usage and gather system info
	totalContainers := systemInfo.Containers
	runningContainers := systemInfo.ContainersRunning
	stoppedContainers := systemInfo.ContainersStopped
	totalImgSize := int64(0)
	ac