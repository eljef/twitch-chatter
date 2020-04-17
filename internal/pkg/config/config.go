package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

// Data is the configuration data holder
type Data struct {
	Plugins DataPlugins       `toml:"plugins"` // Plugins holds the plugins to enable
	Scenes  map[string]string `toml:"scenes"`  // Scenes are the commands to switch scenes
	OBS     DataOBS           `toml:"obs"`     // OBS holds configuration for OBS
	Twitch  DataTwitch        `toml:"twitch"`  // Twitch holds twitch connection data
}

// DataPlugins holds the plugins configuration data
type DataPlugins struct {
	Approved      []string `toml:"approved_users"` // Approved is the list of approved users
	InChannel     bool     `toml:"in_channel"`     // InChannel allows moderators to run commands in chat
	Sceneswitcher bool     `toml:"sceneswitcher"`  // SceneSwitcher is enabled or not
}

// DataOBS holds obs configuration data
type DataOBS struct {
	Host string `toml:"host"`     // Host is the address the OBS websocket is listening on
	Port int    `toml:"port"`     // Port is the port the OBS websocket is listening on
	Pass string `toml:"password"` // Pass is the password for the OBS websocket
}

// DataTwitch holds Twitch configuration data
type DataTwitch struct {
	Channel string `toml:"channel"`     // Channel is the Twitch channel to join
	Name    string `toml:"username"`    // Name is the Twitch username to connect as
	Token   string `toml:"oauth_token"` // Token is the OAUTH Token to use for connections
}

// getConfigPath finds the path to the config file to use
func getConfigPath(paths []string) (string, error) {
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", errors.New("could not find config file")
}

// ReadConfig reads the configuration file for the project
func ReadConfig(paths []string) (*Data, error) {
	configFile, err := getConfigPath(paths)
	if err != nil {
		return nil, err
	}

	ret := &Data{}
	_, err = toml.DecodeFile(configFile, ret)

	return ret, err
}
