// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/drivers/gpio"
	"github.com/rich1111/gobot/platforms/beaglebone"
)

func main() {
	beagleboneAdaptor := beaglebone.NewAdaptor()
	servo := gpio.NewServoDriver(beagleboneAdaptor, "P9_14")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
