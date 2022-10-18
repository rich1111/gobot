/*
Package digispark provides the Gobot adaptor for the Digispark ATTiny-based USB development board.

Installing:

This package requires installing `libusb`.
Then you can install the package with:

	go get -u -d github.com/rich1111/gobot/platforms/digispark

Example:

	package main

	import (
		"time"

		"github.com/rich1111/gobot"
		"github.com/rich1111/gobot/drivers/gpio"
		"github.com/rich1111/gobot/platforms/digispark"
	)

	func main() {
		digisparkAdaptor := digispark.NewAdaptor()
		led := gpio.NewLedDriver(digisparkAdaptor, "0")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{digisparkAdaptor},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For further information refer to digispark README:
https://github.com/hybridgroup/gobot/blob/master/platforms/digispark/README.md
*/
package digispark // import "github.com/rich1111/gobot/platforms/digispark"
