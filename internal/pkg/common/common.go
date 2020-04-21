// Copyright (c) 2020 Jef Oliver. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package common

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

// Time returns a formatted time string
func Time() string {
	now := time.Now()
	return now.Format("2006/01/02 15:04:05")
}

// IsBroadcaster determines if the message came from the broadcaster
func IsBroadcaster(user twitch.User) bool {
	_, ok := user.Badges["broadcaster"]

	return ok
}

// IsModerator determines if the message came from a moderator
func IsModerator(user twitch.User) bool {
	_, ok := user.Badges["moderator"]

	return ok
}

// Log logs a message to the console
func Log(data interface{}) {
	/* #nosec */
	_, _ = fmt.Printf("[chatt] %s %s\n", Time(), data)
}

// LogError logs an error message to the console
func LogError(data interface{}) {
	Log(fmt.Sprintf("ERROR OCCURED: %v", data))
}

// ModCheck determines if the user is a moderator level user
func ModCheck(user twitch.User) bool {
	if IsModerator(user) {
		return true
	}

	return IsBroadcaster(user)
}
