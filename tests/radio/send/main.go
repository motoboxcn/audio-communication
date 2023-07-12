package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(20)
	pin.Output()

	sendCode := 2147483647
	sendData(pin, sendCode)

	fmt.Println("Data sent successfully!")
}

func sendData(pin rpio.Pin, code int) {
	fmt.Printf("Sending data: %v", code)

	for i := 0; i < 10; i++ {
		pin.Toggle()
		time.Sleep(time.Millisecond * 200)
	}

	// 转化为二进制数组
	binaryData := make([]byte, 4)
	binaryData[0] = byte(code >> 24)
	binaryData[1] = byte(code >> 16)
	binaryData[2] = byte(code >> 8)
	binaryData[3] = byte(code)

	for _, data := range binaryData {
		for i := 0; i < 8; i++ {
			if (data & 0x80) == 0x80 {
				pin.High()
			} else {
				pin.Low()
			}
			data <<= 1
			time.Sleep(time.Millisecond * 2)
		}
	}

	pin.Low()
	time.Sleep(time.Millisecond * 100)
}
