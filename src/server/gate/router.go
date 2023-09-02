package gate

import (
	"leafServer/src/server/game"
	"leafServer/src/server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
