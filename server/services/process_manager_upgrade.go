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
	