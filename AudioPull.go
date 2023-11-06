package main

import (
	"fmt"
	pa "github.com/gordonklaus/portaudio"
)

func AudioPull() {
	if err := pa.Initialize(); err != nil {
        fmt.Printf("PortAudio initialization failed: %v\n", err)
        return
    }
    defer pa.Terminate()
	fmt.Println(pa.VersionText())
}