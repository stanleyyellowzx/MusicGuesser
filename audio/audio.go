package audio

import (
    "os"
    "log"
    "time"
	"math/rand/v2"
    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
)

var globalSampleRate = beep.SampleRate(48000)
var initialized = false

func PlayAudioClip(filename string, songLength int) {
	initSpeaker()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
    streamer, format, err := mp3.Decode(f)
    if err != nil {
		log.Fatal(err)
	}
    defer streamer.Close()

	// generate time to skip
	secondsToSkip := generateAudioStart10Seconds(songLength)
	samplesToSkip := format.SampleRate.N(time.Duration(secondsToSkip) * time.Second)
	streamer.Seek(samplesToSkip)

	// resample to match current song
	var stream beep.Streamer = streamer

	if format.SampleRate != globalSampleRate {
        stream = beep.Resample(4, format.SampleRate, globalSampleRate, streamer)
    }

	// play audio for 10 seconds
	samplesToPlay := format.SampleRate.N(10 * time.Second)
	limited := beep.Take(samplesToPlay, stream)

	// play audio
	done := make(chan struct{})
	speaker.Play(beep.Seq(limited, beep.Callback(func() {
		close(done)
	})))

	<-done
}

func initSpeaker() {
	if initialized {
		return
	}
	speaker.Init(globalSampleRate, globalSampleRate.N(time.Second/10))
	initialized = true
}

func generateAudioStart10Seconds(num int) int{
	// the latest point that the audio can start at is 10 seconds before the end of the song
	const startingPoint int = 10
	return rand.IntN(num - startingPoint)
}