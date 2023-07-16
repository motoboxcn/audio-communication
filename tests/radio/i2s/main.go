package main

import (
	"log"

	"github.com/gordonklaus/portaudio"
)

// i2s 语音通信 将语音数据转换为二进制数据
func i2sRecord(out []int32) {
	// 初始化 portaudio
	err := portaudio.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer portaudio.Terminate()
	// 打开输入流
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(out), out)
	// 错误处理
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	// 开始录制语音数据
	err = stream.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	// 循环录制语音数据
	for {
		// 从输入流中读取语音数据到缓冲区中
		err := stream.Read()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	out := make([]int32, 128)
	i2sRecord(out)
}
