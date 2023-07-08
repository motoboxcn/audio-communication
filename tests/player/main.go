package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// apt-get install libasound2-dev

func main() {
	// 打开音频文件
	file, err := os.Open("music.mp3")
	if err != nil {
		fmt.Println("无法打开音频文件:", err)
		return
	}
	defer file.Close()

	// 解码MP3文件
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		fmt.Println("无法解码MP3文件:", err)
		return
	}

	// 创建音频播放设备
	context, err := oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		fmt.Println("无法创建音频设备:", err)
		return
	}
	defer context.Close()

	// 创建音频播放器
	player := context.NewPlayer()
	defer player.Close()

	// 播放音频
	_, err = io.Copy(player, decoder)
	if err != nil {
		fmt.Println("播放音频失败:", err)
		return
	}
}
