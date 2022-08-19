package LED

import "machine"

var (
	light = machine.LED
)

type Mode struct{}

func (Mode) Init() {
	light = machine.LED
	light.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (Mode) Power(mode int) {
	switch mode {
	case 0:
		light.Low()
	case 1:
		light.High()
	default:
		panic("bad number given please use 0 or 1")
	}
}
