package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()
	for {
		chk(stream.Read())
		fmt.Println(in)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
