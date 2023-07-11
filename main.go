package main

import (
	"os"
	"time"

	"github.com/motoboxcn/audio-communication/pkg/audio"
	"github.com/motoboxcn/audio-communication/pkg/bluetooth"
	"github.com/patsnapops/noop/log"
)

const (
	VIMOTOV9S   string = "1C:05:B7:11:D8:EE"
	BTS_06      string = "41:42:2C:8D:52:A3"
	mikasiPhone string = "5C:87:30:3C:29:6B"
	mk          string = "48:74:6E:85:84:D0"
	XiaoMi      string = "00:42:79:2C:3D:D7"
)

func init() {
	log.Default().WithLevel(log.DebugLevel).Init()
}

func main() {
	args := os.Args
	bluetoothClient := bluetooth.NewBluetoothClient(XiaoMi)

	if len(args) < 2 {
		audioStart()
	} else {
		switch args[1] {
		case "con":
			err := bluetoothClient.Connect()
			if err != nil {
				log.Errorf("bluetoothClient.Connect() failed: %v", err)
			}
		case "dis":
			err := bluetoothClient.Disconnect()
			if err != nil {
				log.Errorf("bluetoothClient.Disconnect() failed: %v", err)
			}
		case "audio":
			audioStart()
		default:
			log.Infof("not support args[1]: %v", args[1])
		}
	}
}

func audioStart() {
	audioClient := audio.NewAudioClient()
	audioClient.Start()
	time.Sleep(10 * time.Second)
	audioClient.Stop()
}
