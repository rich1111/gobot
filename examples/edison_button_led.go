// +build example
//
// Do not build by default.

package main

import (
	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/drivers/gpio"
	"github.com/rich1111/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()

	button := gpio.NewButtonDriver(e, "2")
	led := gpio.NewLedDriver(e, "4")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{e},
		[]gobot.Device{led, button},
		work,
	)

	robot.Start()
}
