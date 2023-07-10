package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// apt-get install libasound2-dev
func main() {
	// 打开音频文件夹，循环播放
	cycle := false
	if len(os.Args) != 1 {
		cycle = os.Args[1] == "cycle"
	}

	for {

		filepath.Walk("/root/audio-communication/tests/player", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("无法访问音频文件夹:", err)
				return err
			}
			if info.IsDir() {
				return nil
			}
			if info.Mode().IsRegular() {
				fmt.Println(path)
				// 如果是音乐文件则打开播放
				if filepath.Ext(path) == ".mp3" {
					file, err := os.Open("music.mp3")
					if err != nil {
						fmt.Println("无法打开音频文件:", err)
						return err
					}
					defer file.Close()
					// 解码MP3文件
					decoder, err := mp3.NewDecoder(file)
					if err != nil {
						fmt.Println("无法解码MP3文件:", err)
						return err
					}
					// 创建音频播放设备
					context, err := oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
					if err != nil {
						fmt.Println("无法创建音频设备:", err)
						return err
					}
					defer context.Close()
					// 创建音频播放器
					player := context.NewPlayer()
					defer player.Close()
					// 播放音频
					_, err = io.Copy(player, decoder)
					if err != nil {
						fmt.Println("播放音频失败:", err)
						return err
					}
				}
			}
			return nil
		})

		if !cycle {
			break
		}
	}
}
