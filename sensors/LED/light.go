package LED

import "machine"

var (
	Light = machine.LED
)

type Config struct{}

func (Config) Init() {
	Light = machine.LED
	Light.Configure(machine.PinConfig{Mode: machine.PinOutput})
}
