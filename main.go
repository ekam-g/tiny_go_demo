package main

import (
	"machine"

	"tinygo.org/x/drivers/hcsr04"
	

)
const (

)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	Ultrasonic := machine.GPIO
	Ultrasonic.Configure(machine.GPIConfig{})
	sensor := hcsr04.New(Ultrasonic)
	sensor.Configure()
	for {
		if sensor.ReadDistance() > 20 {
			led.Low()
		} else {
			led.High()
		}
	}
}

// tinygo flash -target arduino main.go
