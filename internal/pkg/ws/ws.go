// Copyright (c) 2020 Jef Oliver. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package ws

import (
	"git.eljef.me/go/twitch-chatter/internal/pkg/common"
	"git.eljef.me/go/twitch-chatter/internal/pkg/config"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

var (
	wsClient obsws.Client
)

// Connect initializes the websocket connection to the OBS WebSocket
func Connect(data config.DataOBS) error {
	wsClient = obsws.Client{Host: data.Host, Port: data.Port, Password: data.Pass}
	return wsClient.Connect()
}

// Disconnect disconnects the websocket connection from the OBS WebSocket
func Disconnect() {
	if err := wsClient.Disconnect(); err != nil {
		common.LogError("%v", err)
	}
}

// Get returns an obsws.Client for usage.
func Get() obsws.Client {
	return wsClient
}
