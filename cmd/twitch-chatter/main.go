package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"git.eljef.me/go/twitch-chatter/internal/pkg/common"
	"git.eljef.me/go/twitch-chatter/internal/pkg/config"
	"git.eljef.me/go/twitch-chatter/internal/pkg/dispatcher"
	"git.eljef.me/go/twitch-chatter/internal/pkg/sceneswitch"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/gempir/go-twitch-irc/v2"
)

func badExit(data interface{}) {
	common.LogError(data)
	os.Exit(-1)
}

func configPlugins(configData *config.Data) {
	dispatcher.Config(configData)
	sceneswitch.Config(configData)
}

func main() {
	usr, err := user.Current()
	if err != nil {
		badExit(err)
	}

	checkPaths := []string{
		"twitch-chatter.toml",
		path.Join(usr.HomeDir, "twitch-chatter.toml"),
		"/etc/twitch-chatter.toml",
	}

	configData, err := config.ReadConfig(checkPaths)
	if err != nil {
		badExit(err)
	}

	common.WSClient = obsws.Client{Host: configData.OBS.Host, Port: configData.OBS.Port,
		Password: configData.OBS.Pass}
	if err := common.WSClient.Connect(); err != nil {
		badExit(err)
	}
	defer common.WSClient.Disconnect()

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
