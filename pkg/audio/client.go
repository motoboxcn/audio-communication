package audio

/*
PortAudio是一个跨平台音频I/O库，它不提供操作系统级别的音频设备管理功能。
然而，你可以使用PortAudio的API在程序中选择特定的音频输入设备。
对于需要使用蓝牙设备作为音频输入的应用程序，你首先需要获取所有音频设备的列表，
然后选择蓝牙设备作为输入设备。此过程涉及到PortAudio API的使用，具体请参考PortAudio的官方文档。
*/

import (
	"github.com/gordonklaus/portaudio"
	"github.com/patsnapops/noop/log"
)

type AudioClient struct {
	Stream                   *portaudio.Stream
	InputDevices, OutDevices portaudio.StreamDeviceParameters
}

func NewAudioClient() *AudioClient {
	err := portaudio.Initialize()
	if err != nil {
		log.Panicf("portaudio.Initialize() failed: %v", err)
	}
	return &AudioClient{
		InputDevices: portaudio.StreamDeviceParameters{
			Device: GetInputDevices("pulse"),
		},
		OutDevices: portaudio.StreamDeviceParameters{
			Device: nil,
		},
	}
}

// 获取输入设备列表
func GetInputDevices(name string) (device *portaudio.DeviceInfo) {
	devices, err := portaudio.Devices()
	if err != nil {
		log.Panicf("portaudio.Devices() failed: %v", err)
	}
	for _, device := range devices {
		log.Infof("device: %v", device)
		if device.Name == name {
			return device
		}
	}
	return nil
}

func (c *AudioClient) Start() {
	in := make([]int32, 64)
	// 非阻塞
	stream, err := portaudio.OpenStream(portaudio.StreamParameters{
		Input:  c.InputDevices,
		Output: c.OutDevices,
	}, in, len(in))
	if err != nil {
		log.Panicf("portaudio.OpenDefaultStream() failed: %v", err)
	}
	c.Stream = stream
	err = c.Stream.Start()
	if err != nil {
		log.Panicf("c.stream.Start() failed: %v", err)
	}
	go func() {
		for {
			err := c.Stream.Read()
			if err != nil {
				return
			}
		}
	}()
}

func (c *AudioClient) Stop() {
	err := c.Stream.Stop()
	if err != nil {
		log.Panicf("c.stream.Stop() failed: %v", err)
	}
}
