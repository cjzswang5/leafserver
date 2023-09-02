package internal

import (
	"reflect"

	"github.com/cjzswang5/leafServer/src/server/msg"

	"github.com/cjzswang5/leaf/gate"
	"github.com/cjzswang5/leaf/log"
)

func init() {
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	log.Debug("hello %v", m.Name)
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}
