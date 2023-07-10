package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func listBluetoothDevices() {
	out, err := exec.Command("bluetoothctl", "devices").Output()
	if err != nil {
		log.Fatal("Failed to execute command: bluetoothctl devices")
		return
	}
	fmt.Println(string(out))
}

// scanBluetoothDevices scans for bluetooth devices
func scanBluetoothDevices() {
	out, err := exec.Command("bluetoothctl", "scan", "on").Output()
	if err != nil {
		log.Fatal("Failed to execute command: bluetoothctl scan on")
		return
	}
	fmt.Println(string(out))
}

// connectBluetoothDevice connects to a bluetooth device
func connectBluetoothDevice(mac string) {
	out, err := exec.Command("bluetoothctl", "connect", mac).Output()
	if err != nil {
		log.Fatalf("Failed to execute command: bluetoothctl connect %s", mac)
		return
	}
	fmt.Println(string(out))
}

// disconnectBluetoothDevice disconnects from a bluetooth device
func disconnectBluetoothDevice(mac string) {
	out, err := exec.Command("bluetoothctl", "disconnect", mac).Output()
	if err != nil {
		log.Fatalf("Failed to execute command: bluetoothctl disconnect %s", mac)
		return
	}
	fmt.Println(string(out))
}

// pairBluetoothDevice pairs with a bluetooth device
func pairBluetoothDevice(mac string) {
	out, err := exec.Command("bluetoothctl", "pair", mac).Output()
	if err != nil {
		log.Fatalf("Failed to execute command: bluetoothctl pair %s", mac)
		return
	}
	fmt.Println(string(out))
	// trust
	out, err = exec.Command("bluetoothctl", "trust", mac).Output()
	if err != nil {
		log.Fatalf("Failed to execute command: bluetoothctl trust %s", mac)
		return
	}
	fmt.Println(string(out))
}

// systemctl stop pulseaudio
func stopPulseAudio() {
	out, err := exec.Command("systemctl", "stop", "pulseaudio").Output()
	if err != nil {
		log.Fatal("Failed to execute command: systemctl stop pulseaudio")
		return
	}
	fmt.Println(string(out))
}

// start pulseaudio
func startPulseAudio() {
	out, err := exec.Command("systemctl", "start", "pulseaudio").Output()
	if err != nil {
		log.Fatal("Failed to execute command: systemctl start pulseaudio")
		return
	}
	fmt.Println(string(out))
}

func main() {
	args := os.Args
	// mac := "41:42:2C:8D:52:A3"
	mac := "00:42:79:2C:3D:D7" //xiaomi
	switch args[1] {
	case "connect":
		startPulseAudio()
		connectBluetoothDevice(mac)
	case "disconnect":
		disconnectBluetoothDevice(mac)
	case "scan":
		scanBluetoothDevices()
	case "list":
		listBluetoothDevices()
	case "pair":
		pairBluetoothDevice(mac)
	case "stop":
		disconnectBluetoothDevice(mac)
		stopPulseAudio()
	case "restart":
		stopPulseAudio()
		startPulseAudio()
	default:
		fmt.Println("Invalid command")
	}
}
