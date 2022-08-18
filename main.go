package main

import (
	"machine"
	"time"
	"tiny_go/sensors/LED"
	"tiny_go/sensors/ultraSonic"
)

func main() {
	LED.Config{}.Init()
	device := ultraSonic.New(machine.Pin(2), machine.Pin(3))
	sen := &device
	sen.Configure()
	for {
		if sen.ReadDistance() > 10 {
			LED.Light.High()
		} else {
			LED.Light.Low()
		}
		time.Sleep(time.Second * 2)
	}
}
