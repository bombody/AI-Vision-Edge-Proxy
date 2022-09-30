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
	"os"
	"strconv"
	"testing"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	badger "github.com/dgraph-io/badger/v2"
)

var (
	testDBPath = "/tmp/chrystest"
)

func setupDB() (*badger.DB, error) {
	if _, err := os.Stat(testDBPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		errDir := os.Mkdir(testDBPath, 0755)
		if errDir != nil {
			g.Log.Error("failed to create directiory for DB", testDBPath, errDir)
			return nil, errDir
		}
	} else {
		err := os.RemoveAll(testDBPath)
		if err != nil {
			return ni