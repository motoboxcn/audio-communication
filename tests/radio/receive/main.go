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
	pin.Input()

	data := make([]byte, 4)

	for {
		for i := 0; i < len(data); i++ {
			for j := 0; j < 8; j++ {
				time.Sleep(time.Millisecond * 1)
				data[i] <<= 1
				data[i] |= uint8(pin.Read())
			}
		}

		parseData(data)

		time.Sleep(time.Millisecond * 500)
	}
}

func parseData(data []byte) {
	value := uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	fmt.Println("Received data:", value)
}
