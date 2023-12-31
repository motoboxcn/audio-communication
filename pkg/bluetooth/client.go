package bluetooth

import (
	"os/exec"

	"github.com/patsnapops/noop/log"
)

type BluetoothClient struct {
	Mac string
}

func NewBluetoothClient(mac string) *BluetoothClient {
	err := startPulseAudio()
	if err != nil {
		log.Fatalf("startPulseAudio() failed: %v", err)
	}
	return &BluetoothClient{
		Mac: mac,
	}
}

// pair pairs with a bluetooth device
func (c *BluetoothClient) Pair() error {
	_, err := exec.Command("bluetoothctl", "pair", c.Mac).Output()
	if err != nil {
		return err
	}
	// trust
	_, err = exec.Command("bluetoothctl", "trust", c.Mac).Output()
	if err != nil {
		return err
	}
	return nil
}

// trust to a bluetooth device
func (c *BluetoothClient) Trust() error {
	_, err := exec.Command("bluetoothctl", "trust", c.Mac).Output()
	if err != nil {
		return err
	}
	return nil
}

// connect connects to a bluetooth device
func (c *BluetoothClient) Connect() error {
	out, err := exec.Command("bluetoothctl", "connect", c.Mac).Output()
	if err != nil {
		log.Infof(string(out))
		return err
	}
	return nil
}

// disconnect disconnects from a bluetooth device
func (c *BluetoothClient) Disconnect() error {
	stopPulseAudio()
	_, err := exec.Command("bluetoothctl", "disconnect", c.Mac).Output()
	if err != nil {
		return err
	}
	return nil
}

func stopPulseAudio() error {
	out, err := exec.Command("pulseaudio", "--kill").Output()
	if err != nil {
		log.Infof("stopPulseAudio failed: %v", out)
		return err
	}
	return nil
}

func startPulseAudio() error {
	out, err := exec.Command("pulseaudio", "--start").Output()
	if err != nil {
		log.Infof("startPulseAudio failed: %v", out)
		return err
	}
	return nil
}
