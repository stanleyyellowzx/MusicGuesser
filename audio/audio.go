package audio

import (
    "os"
    "log"
    "time"
    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
)

func PlayAudio(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
    streamer, format, err := mp3.Decode(f)
    if err != nil {
		log.Fatal(err)
	}
    defer streamer.Close()
    
    //sr := format.SampleRate * 2
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    //resampled := beep.Resample(4, format.SampleRate, sr, streamer)
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	<-done
}

func PlayAudioClip(filename string) {
	
}