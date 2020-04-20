// Copyright (c) 2020 Jef Oliver. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package sceneswitch

import (
	"fmt"

	"git.eljef.me/go/twitch-chatter/internal/pkg/common"
	"git.eljef.me/go/twitch-chatter/internal/pkg/config"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

var (
	enabled bool
	scenes  map[string]string
)

// Config configures the module
func Config(data *config.Data) {
	enabled = data.Plugins.Sceneswitcher
	scenes = data.Scenes
}

// Handle handles switch of OBS scenes
func Handle(sceneName string, username string) {
	if !enabled {
		return
	}

	setScene(sceneName, username)
}

// logMessage logs the request
func logMessage(sceneName string, username string) {
	common.Log(fmt.Sprintf("**** Scene Switch: %s -> Requested By: %s ****", sceneName, username))
}

// setScene sets the actual scene in OBS
func setScene(sceneName string, username string) {
	if _, ok := scenes[sceneName]; ok {
		logMessage(sceneName, username)
		req := obsws.NewSetCurrentSceneRequest(scenes[sceneName])
		if err := req.Send(common.WSClient); err != nil {
			common.LogError(err)
		}
	}
}
