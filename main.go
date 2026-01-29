package main

import (
	"fmt"
	"github.com/stanleyyellowzx/MusicGuesser/audio"
	"github.com/stanleyyellowzx/MusicGuesser/database"
)

func main() {
	var userInput int
	filename := "audio_files/"
	database.ConnectToDatabase()

	songData, err := database.GetAllSongData()
	if (err != nil) {
		return
	}
	fmt.Println("Which song would you like to play?")
	printSlice(songData)

	fmt.Scanln(&userInput)
	playSong := fmt.Sprintf("%s%s", filename, songData[userInput].Song_file_name)
	fmt.Println("Playing song: ", songData[userInput].Song_name)
	audio.PlayAudio(playSong)

	fmt.Println("Playing audio")
}

func printSlice(slice []database.SongData) {
	for index, element := range slice {
		fmt.Println("Index: ", index, ", Element: ", element)
	}
}