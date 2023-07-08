package main

import (
	"context"
	"fmt"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/patsnapops/noop/log"
)

// 1C:05:B7:11:D8:EE VIMOTO V9S
// 41:42:2C:8D:52:A3 BTS-06
// 5C:87:30:3C:29:6B       mikas iPhone

var deviceName string = "BTS-06"

func init() {
	// 初始化蓝牙适配器
	log.Default().WithLevel(log.DebugLevel).Init()
	d, err := dev.NewDevice("default")
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)
}

// scan bluetooth devices
func scan() {
	// 扫描蓝牙设备
	err := ble.Scan(
		context.Background(),
		false, // 允许重复扫描
		func(a ble.Advertisement) {
			fmt.Println(a.Addr(), a.LocalName())
			if a.LocalName() == deviceName {
				fmt.Println("Found device")
				connect()
			}
		}, // 处理器
		func(a ble.Advertisement) bool {
			return a.LocalName() == ""
		}, // 过滤器
	)
	if err != nil {
		log.Fatalf("Failed to scan: %s", err)
	}

	fmt.Println("Scan completed")
}

// connect bluetooth device
func connect() {
	var f ble.AdvFilter = func(a ble.Advertisement) bool {
		return a.LocalName() == deviceName
	}
	p, err := ble.Connect(
		context.Background(),
		f,
	)
	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	fmt.Printf("Connected to %s\n", p.Addr())
}

func main() {
	scan()
	// connect()
}
