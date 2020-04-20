# Twitch Chatter

Twitch Chatter is a small utility program that integrates with
[OBS](https://obsproject.com/) via the
[OBS WebSocket](https://github.com/Palakis/obs-websocket) plugin. Twitch Chatter
allows for trusted users to whisper commands to the broadcaster to control OBS.
Additionally, Twitch Chatter can be configured to listen for commands issued by
moderators in your channel.

## Requirements

The OBS WebSocket plugin must be installed and configured in OBS. After starting
OBS, you can then start Twitch Chatter.

## Configuration

Twitch Chatter requires a filled out configuration file. An example exists in
the configs folder. Once this is filled out, the file should be placed in one of
the following locations.

* Same folder as the executable for Twitch Chatter
* Users home directory
  * ~/.twitch-chatter.toml on Linux/Unix,
  * %USERPROFILE%/twitch-chatter.toml on Windows
* The system directory
  * /etc/twitch-chatter.toml on Linux / Unix
  * %SYSTEMROOT%/twitch-chatter.toml on Windows

### OAUTH

Twitch Chatter uses an OAUTH token to connect to the Twitch IRC (Internet Relay
Chat) interface. An OAUTH token can be obtained from the
[Twitch Chat OAUTH Password Generator](https://twitchapps.com/tmi/). This token
must be saved in the configuration file.

## Running

Once the configuration file has been saved, and OBS has been started with the
WebSocket plugin installed, simply run Twitch Chatter.

## Plugins

### Sceneswitcher

Sceneswitcher allows an authorized user to whisper the `!scene` command to the
broadcaster to switch OBS scenes.

```!scene scenename```

## Installing

* Download one of the
  [releases](https://git.eljef.me/go/twitch-chatter/releases)
* go get git.eljef.me/go/twitch-chatter/cmd/twitch-chatter
  * This will install twitch-chatter in your $GOHOME/bin folder
