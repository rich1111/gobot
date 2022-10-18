// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_lidarlite.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/drivers/i2c"
	"github.com/rich1111/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	lidar := i2c.NewLIDARLiteDriver(firmataAdaptor)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			distance, _ := lidar.Distance()
			fmt.Println("Distance", distance)
		})
	}

	robot := gobot.NewRobot("lidarbot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{lidar},
		work,
	)

	robot.Start()
}
