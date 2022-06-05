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

package models

const (
	PrefixSettingsKey               = "/settings/"
	PrefixSettingsDockerTagVersions = "/dockertagsettings/"

	SettingDefaultKey = "default"
)

// Settings - keeping setting on the edge
type Settings struct {
	Name          string `json:"name"`                      // name of the setting
	EdgeKey       string `json:"edge_key,omitempty"`        // edge key genera