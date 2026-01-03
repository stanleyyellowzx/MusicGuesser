package database

import (
	"testing"
	"fmt"
)

func TestNoConnectionToDatabase(t *testing.T) {
	fmt.Println("test.go file")
	fmt.Println("testing GetAllSongData before connection to database")
	GetAllSongData()
	fmt.Println("testing GetArtistSongData before connection to database")
	GetArtistSongData("Yorushika")
}

func TestConnectionToDatabase(t *testing.T) {
	fmt.Println("connecting to database")
	ConnectToDatabase()
	fmt.Println("testing GetAllSongData after connection to database")
	GetAllSongData()
	fmt.Println("testing GetArtistSongData after connection to database")
	GetArtistSongData("Yorushika")
	CloseDatabase()
}