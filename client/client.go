package client

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"bufio"
	"os"
	"github.com/stanleyyellowzx/MusicGuesser/audio"
	"github.com/stanleyyellowzx/MusicGuesser/database"
)

func Start() {
	var userInput string
	audio_directory := "audio_files/"
	database.ConnectToDatabase()
	reader := bufio.NewReader(os.Stdin)

	songData, err := database.GetAllSongData()
	if (err != nil) {
		return
	}

	fmt.Println("Welcome to MusicGuesser")
	fmt.Println("Each game has 5 songs randomly selected from the selection of songs")
	fmt.Println("Press enter to start")
	fmt.Scanln()
	for {
		score := 0
		selectedSongs := generateRandomSongs(5)
		fmt.Println("Playing Songs")
		for _, song := range selectedSongs {
			playSong := fmt.Sprintf("%s%s", audio_directory, songData[song].Song_file_name)
			audio.PlayAudioClip(playSong, songData[song].Duration)

			fmt.Println("Song name?")
			input, _ := reader.ReadString('\n')
			userInput = strings.TrimSpace(input)
			if strings.EqualFold(userInput, songData[song].Song_name) {
				score++
			}

			// flush out userInput
			userInput = ""
		}

		fmt.Println("All songs completed")
		fmt.Println("Final score is: ", score)
		fmt.Println("Enter Q to quit or any other key to play again")
		fmt.Scanln(&userInput)
		if strings.EqualFold(userInput, "q") {
			break
		}
	}
}

func printTest(slice []int) {
	for index, element := range slice {
		fmt.Println(index, ":", element)
	}
}

func generateRandomSongs(numSongs int) []int{
	songs := rand.Perm(numSongs)
	// add 1 to each element for right now since there is not audio file for Nautilus by Yorushika
	for index, element := range songs {
		songs[index] = element + 1
	}
	return songs
}