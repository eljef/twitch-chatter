// Copyright (c) 2020 Jef Oliver. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package dispatcher

import (
	"strings"

	"git.eljef.me/go/twitch-chatter/internal/pkg/common"
	"git.eljef.me/go/twitch-chatter/internal/pkg/config"
	"git.eljef.me/go/twitch-chatter/internal/pkg/sceneswitch"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	configData []string
)

// Config configures the module
func Config(data *config.Data) {
	configData = data.Plugins.Approved
}

// DispatchChannel dispatches a command requested in the channel
func DispatchChannel(message twitch.PrivateMessage) {
	if !strings.HasPrefix(message.Message, "!") {
		return
	}

	commandString := strings.SplitN(message.Message, " ", 2)
	if commandString[0] == "!scene" && common.ModCheck(message.User) {
		sceneswitch.Handle(commandString[1], message.User.Name)
	}
}

// DispatchWhisper dispatches a command that matches a plugin command
func DispatchWhisper(message twitch.WhisperMessage) {
	if !strings.HasPrefix(message.Message, "!") {
		return
	}

	for _, username := range configData {
		if strings.EqualFold(username, message.User.Name) {
			goto dispatch
		}
	}

	return

dispatch:
	commandString := strings.SplitN(message.Message, " ", 2)
	if commandString[0] == "!scene" {
		sceneswitch.Handle(commandString[1], message.User.Name)
	}
}
