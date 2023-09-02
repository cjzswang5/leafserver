package main

import (
	"github.com/cjzswang5/leafServer/src/server/conf"
	"github.com/cjzswang5/leafServer/src/server/game"
	"github.com/cjzswang5/leafServer/src/server/gate"
	"github.com/cjzswang5/leafServer/src/server/login"

	"github.com/cjzswang5/leaf"
	lconf "github.com/cjzswang5/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
