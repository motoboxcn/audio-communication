package main

import (
	"time"

	"github.com/motoboxcn/audio-communication/pkg/audio"
	"github.com/patsnapops/noop/log"
)

func init() {
	log.Default().WithLevel(log.DebugLevel).Init()
}

func main() {
	// bluetoothClient := bluetooth.NewBluetoothClient("1C:05:B7:11:D8:EE")
	// err := bluetoothClient.Connect()
	// if err != nil {
	// 	log.Errorf("bluetoothClient.Connect() failed: %v", err)
	// }

	// audio
	audioClient := audio.NewAudioClient()
	audioClient.Start()
	time.Sleep(10 * time.Second)
	audioClient.Stop()
}
