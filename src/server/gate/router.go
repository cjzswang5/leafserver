package gate

import (
	"github.com/cjzswang5/leafServer/src/server/game"
	"github.com/cjzswang5/leafServer/src/server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
