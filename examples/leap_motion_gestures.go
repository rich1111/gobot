// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/platforms/leap"
)

func main() {
	leapMotionAdaptor := leap.NewAdaptor("127.0.0.1:6437")
	l := leap.NewDriver(leapMotionAdaptor)

	work := func() {
		l.On(leap.GestureEvent, func(data interface{}) {
			printGesture(data.(leap.Gesture))
		})
	}

	robot := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	robot.Start()
}

func printGesture(gesture leap.Gesture) {
	fmt.Println("Gesture", gesture)
}
