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
	activeImages := 0
	totalVolumeSize := int64(0)
	activeVolumes := int64(0)

	for _, im := range diskUsage.Images {
		activeImages += int(im.Containers)
		totalImgSize += im.SharedSize
	}
	for _, v := range diskUsage.Volumes {
		activeVolumes += v.UsageData.RefCount
		totalVolumeSize += v.UsageData.Size
	}

	stats.Containers = totalContainers
	stats.ContainersRunning = runningContainers
	stats.ContainersStopped = stoppedContainers
	stats.ActiveImages = int(activeImages)
	stats.TotalVolumeSize = totalVolumeSize
	stats.TotalActiveVolumes = int(activeVolumes)
	stats.GatewayID = sett.GatewayID
	stats.TotalImageSize = totalImgSize

	stats.ContainersStats = make([]*models.ProcessStats, 0)

	pList, err := pm.List()
	if err != nil {
		g.Log.Error("failed to list all containers", err)
		return nil, err
	}

	// gather all container stats
	for _, process := range pList {
		c, err := cl.ContainerGet(process.ContainerID)
		if err != nil {
			g.Log.Error("failed to get container from docker system", err)
			continue
		}
		n := c.Name
		// skip default running components
		if strings.Contain