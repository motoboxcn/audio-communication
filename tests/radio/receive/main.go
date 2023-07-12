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
	for {
		receiveData := receiveCode(pin)
		fmt.Printf("Received data: %v\n", receiveData)
		time.Sleep(time.Millisecond * 100)
	}

}

func receiveCode(pin rpio.Pin) int {
	code := 0

	for i := 0; i < 32; i++ {
		code <<= 1
		if pin.Read() == rpio.High {
			code |= 1
		}
		time.Sleep(time.Millisecond * 2)
	}

	return code
}
