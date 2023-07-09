package main

import (
	"context"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/patsnapops/noop/log"
)

func main() {
	log.Default().WithLevel(log.DebugLevel).Init()
	d, err := dev.NewDevice("default")
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// 设置超时时间
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 1*time.Minute))

	log.Infof("开始扫描...")
	err = ble.Scan(ctx, false, advHandler, nil)
	if err != nil {
		log.Fatalf("can't scan: %s", err)
	}
}

func advHandler(a ble.Advertisement) {
	log.Infof("发现设备: %s %s", a.Addr().String(), a.LocalName())
	// if a.Addr().String() == "dc:05:b7:11:d8:ee" {
	// 	// connect
	// 	log.Infof("开始连接...")
	// 	p, err := ble.Connect(context.Background(), nil)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	log.Infof("连接成功: %s", p.Addr().String())
	// }
}
