package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
	"github.com/tarm/serial"
)

func main() {
	// 初始化RF模块的串口连接
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Failed to open serial port:", err)
		os.Exit(1)
	}

	// 初始化GPIO库
	err = rpio.Open()
	if err != nil {
		fmt.Println("Failed to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// 设置发送引脚为输出模式
	pin := rpio.Pin(4)
	pin.Output()

	// 发送无线电信号
	for {
		// 设置发送引脚为高电平，发射信号
		pin.High()

		// 发送数据到RF模块
		_, err := s.Write([]byte("Hello World"))
		if err != nil {
			fmt.Println("Failed to write to serial port:", err)
		}
		fmt.Printf("Send data to RF module: %s", "Hello World")
		// 等待一段时间
		time.Sleep(time.Second)

		// 设置发送引脚为低电平，停止信号发射
		pin.Low()

		// 读取串口接收到的数据
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println("Failed to read from serial port:", err)
		}
		received := string(buf[:n])
		fmt.Printf("Received data from RF module: %s", received)

		// 等待一段时间
		time.Sleep(time.Second)
	}
}
