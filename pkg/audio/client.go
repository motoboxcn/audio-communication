package audio

import (
	"github.com/gordonklaus/portaudio"
	"github.com/patsnapops/noop/log"
)

type AudioClient struct {
	Stream *portaudio.Stream
	In     []int32
}

func NewAudioClient() *AudioClient {
	err := portaudio.Initialize()
	if err != nil {
		log.Panicf("portaudio.Initialize() failed: %v", err)
	}
	in := make([]int32, 64)
	// 打开默认的输入设备
	inputParams, err := portaudio.DefaultInputDevice()
	if err != nil {
		log.Panicf("portaudio.DefaultInputDevice() failed: %v", err)
	}

	// 打开默认的输出设备
	outputParams, err := portaudio.DefaultOutputDevice()
	if err != nil {
		log.Panicf("portaudio.DefaultOutputDevice() failed: %v", err)
	}

	stream, err := portaudio.OpenStream(portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device: inputParams,
		},
		Output: portaudio.StreamDeviceParameters{
			Device: outputParams,
		},
		FramesPerBuffer: len(in),
		SampleRate:      44100,
		Flags:           portaudio.NoFlag,
	})
	if err != nil {
		log.Panicf("portaudio.OpenDefaultStream() failed: %v", err)
	}

	return &AudioClient{
		Stream: stream,
		In:     in,
	}
}

func (c *AudioClient) Start() {

	err := c.Stream.Start()
	if err != nil {
		log.Panicf("c.stream.Start() failed: %v", err)
	}
	for {
		err := c.Stream.Read()
		if err != nil {
			log.Panicf("c.stream.Read() failed: %v", err)
		}
	}
}

func (c *AudioClient) Stop() {
	err := c.Stream.Stop()
	if err != nil {
		log.Panicf("c.stream.Stop() failed: %v", err)
	}
}
