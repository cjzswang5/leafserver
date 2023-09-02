package msg

import (
	"github.com/cjzswang5/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
}

type Hello struct {
	Name string
}
