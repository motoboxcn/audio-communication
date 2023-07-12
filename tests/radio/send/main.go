package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	// 初始化GPIO库
	err := rpio.Open()
	if err != nil {
		fmt.Println("Failed to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// 设置发送引脚为输出模式
	pin := rpio.Pin(22)
	pin.Output()

	// 发送无线电信号
	for i := 0; i < 10; i++ {
		pin.Low()
		time.Sleep(time.Second)
		pin.High()
		time.Sleep(time.Millisecond * 200)
		pin.Low()
		time.Sleep(time.Millisecond * 100)
		pin.High()
		time.Sleep(time.Millisecond * 200)
	}
	defer pin.Low()
}
