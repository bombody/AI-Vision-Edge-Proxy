package services

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/chryscloud/go-microkit-plugins/docker"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/dgraph-io/badger/v2"
	"github.com/docker/docker/api/types"
	"github.com/hashicorp/go-version"
)

func (pm *ProcessManager) FindUpgrades(imageUpgrade *models.ImageUpgrade) ([]*models.StreamProcess, error) {
	processes, err := pm.List()

	if err != nil {
		g.Log.Error("failed to list local processes", err)
		return nil, err
	}

	upgradesAvailable := make([]*models.StreamProcess, 0)

	if imageUpgrade.CurrentVersion == "" {
		return upgradesAvailable, nil
	}

	currentVersion, vErr := version.NewVersion(imageUpgrade.CurrentVersion)
	if vErr != nil {
		g.Log.Error("version conversion failed", imageUpgrade.CurrentVersion, vErr)
		return nil, vErr
	}

	for _, proc := range processes {
		imgTag := proc.ImageTag
		splitted := strings.Split(imgTag, ":")
		if len(splitted) == 2 {
			ver := splitted[1]
			processVersion, pErr := version.NewVersion(ver)
			if pErr != nil {
				g.Log.Warn("failed to convert version for", ver, pErr)
				continue
			}
			// check if upgrade available
			if currentVersion.GreaterThan(processVersion) {
				proc.UpgradeAvailable = true
				proc.NewerVersion = currentVersion.Original()
				upgradesAvailable = append(upgradesAvailable, proc)
			} else {
				upgradesAvailable = append(upgradesAvailable, proc)
			}
		} else {
			upgradesAvailable = append(upgradesAvailable, proc)
		}
	}

	return upgradesAvailable, nil
}

func (pm *ProcessManager) UpgradeRunningContainer(process *models.StreamProcess, newImage string) (*models.StreamProcess, error) {
	cl := docker.NewSocketClient(docker.Log(g.Log), docker.Host("unix:///var/run/docker.sock"))

	// find container
	containers, err := cl.ContainersList()
	if err != nil {
		g.Log.Error("failed to list running con