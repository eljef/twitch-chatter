// Copyright (c) 2020 Jef Oliver. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"

	"git.eljef.me/go/twitch-chatter/internal/pkg/common"
	"git.eljef.me/go/twitch-chatter/internal/pkg/config"
	"git.eljef.me/go/twitch-chatter/internal/pkg/dispatcher"
	"git.eljef.me/go/twitch-chatter/internal/pkg/sceneswitch"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/gempir/go-twitch-irc/v2"
)

const (
	configFileName = "twitch-chatter.toml"
	version        = "0.0.2"
)

func badExit(data interface{}) {
	common.LogError(data)
	os.Exit(-1)
}

func closeWS() {
	if err := common.WSClient.Disconnect(); err != nil {
		common.LogError(err)
	}
}

func configPlugins(configData *config.Data) {
	dispatcher.Config(configData)
	sceneswitch.Config(configData)
}

func getPaths() []string {
	ret := []string{configFileName}

	if runtime.GOOS == "windows" {
		ret = append(ret, path.Join(os.Getenv("USERPROFILE"), configFileName),
			path.Join(os.Getenv("SYSTEMROOT"), configFileName))
	} else {
		usr, err := user.Current()
		if err != nil {
			badExit(err)
		}
		ret = append(ret, path.Join(usr.HomeDir, fmt.Sprintf(".%s", configFileName)),
			path.Join("/etc", configFileName))
	}

	return ret
}

func main() {
	common.Log(fmt.Sprintf("staring twitch chatter %s", version))
	configData, err := config.ReadConfig(getPaths())
	if err != nil {
		badExit(err)
	}

	common.WSClient = obsws.Client{Host: configData.OBS.Host, Port: configData.OBS.Port,
		Password: configData.OBS.Pass}
	if err := common.WSClient.Connect(); err != nil {
		badExit(err)
	}
	defer closeWS()

	configPlugins(configData)

	common.Log("connecting to twitch")
	client := twitch.NewClient(configData.Twitch.Name, configData.Twitch.Token)
	if configData.Plugins.InChannel {
		client.OnPrivateMessage(dispatcher.DispatchChannel)
	}
	client.OnWhisperMessage(dispatcher.DispatchWhisper)
	client.Join(configData.Twitch.Channel)
	common.Log(fmt.Sprintf("connected to twitch as %s", configData.Twitch.Name))
	err = client.Connect()
	if err != nil {
		badExit(err)
	}
}
