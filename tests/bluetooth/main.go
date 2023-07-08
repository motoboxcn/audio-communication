package main

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/patsnapops/noop/log"
)

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
	var h ble.AdvHandler = func(a ble.Advertisement) {
		fmt.Printf(tea.Prettify(a))
	}

	// 扫描蓝牙设备
	err := ble.Scan(
		context.Background(),
		false, // 允许重复扫描
		h,     // 处理器
		nil,   // 过滤器
	)
	if err != nil {
		log.Fatalf("Failed to scan: %s", err)
	}

	fmt.Println("Scan completed")
}

// connect bluetooth device
func connect() {
	var f ble.AdvFilter = func(a ble.Advertisement) bool {
		return a.LocalName() == "MyDevice"
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
