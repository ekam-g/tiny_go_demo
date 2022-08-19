package main

import (
	"machine"
	"time"
	"tiny_go_demo/sensors/LED"
	"tiny_go_demo/sensors/ultraSonic"
)

func main() {
	led := LED.Mode{}
	led.Init()
	device := ultraSonic.New(machine.Pin(2), machine.Pin(3))
	sen := &device
	sen.Configure()
	for {
		if sen.ReadDistance() > 60 {
			led.Power(1)
		} else {
			led.Power(0)
		}
		if sen.ReadDistance() != 0 {
			println(sen.ReadDistance())
		}
		time.Sleep(time.Second * 2)

	}
}

// tinygo flash -target arduino main.go
