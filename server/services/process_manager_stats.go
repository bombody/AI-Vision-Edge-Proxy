package services

import (
	"strings"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
)

// StatsAllProcesses created a statistics object or all running containers (related to edge)
func (pm *ProcessManager) StatsAllProcesses(sett *models.Settings) (*models.AllStreamProcessStats, e