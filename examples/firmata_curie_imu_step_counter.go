// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_curie_imu_step_counter.go /dev/ttyACM0
*/

package main

import (
	"log"
	"os"
	"time"

	"github.com/rich1111/gobot"
	"github.com/rich1111/gobot/drivers/gpio"
	"github.com/rich1111/gobot/platforms/firmata"
	"github.com/rich1111/gobot/platforms/intel-iot/curie"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	imu := curie.NewIMUDriver(firmataAdaptor)

	work := func() {
		imu.On("Steps", func(data interface{}) {
			log.Println("Steps", data)
		})

		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})

		imu.EnableStepCounter(true)
	}

	robot := gobot.NewRobot("curieBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{imu, led},
		work,
	)

	robot.Start()
}
