// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/drivers/gpio"
	"github.com/rich1111/gobot/platforms/chip"
)

func main() {
	dragonAdaptor := chip.NewAdaptor()
	button := gpio.NewButtonDriver(dragonAdaptor, "GPIO_A")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{dragonAdaptor},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
}
