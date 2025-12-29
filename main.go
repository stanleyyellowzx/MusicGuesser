package main

import (
	"fmt"
	"github.com/stanleyyellowzx/MusicGuesser/audio"
)

func main() {
	filename := "audio_files/eces_boses.mp4"
	audio.PlayAudio(filename)

	fmt.Println("Playing audio")
}