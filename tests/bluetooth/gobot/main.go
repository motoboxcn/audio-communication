package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(
		// "41:42:2C:8D:52:A3",
		"6a:0c:b6:38:89:95",
	)
	battery := ble.NewBatteryDriver(bleAdaptor)

	work := func() {
		gobot.Every(5*time.Second, func() {
			fmt.Println("Battery level:", battery.GetBatteryLevel())
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{battery},
		work,
	)

	robot.Start()
}
