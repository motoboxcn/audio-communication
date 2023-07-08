package main

import (
	"fmt"
	"os"

	"github.com/paypal/gatt"
)

func main() {
	d, err := gatt.NewDevice(gatt.LnxMaxConnections(1))
	if err != nil {
		fmt.Printf("Failed to open device, error: %s", err)
		os.Exit(1)
	}

	d.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
	)

	// 开始扫描蓝牙设备
	d.Init(onStateChanged)
	select {}
}

// 蓝牙设备被发现时的回调函数
func onPeripheralDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	fmt.Printf("Peripheral ID: %s, name: %s, RSSI: %d dB", p.ID(), p.Name(), rssi)
}

// Gatt状态改变时的回调函数
func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Printf("State changed to: %s", s)

	switch s {
	case gatt.StatePoweredOn:
		// 开始扫描蓝牙设备
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}
