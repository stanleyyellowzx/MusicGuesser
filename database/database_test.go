package database

import (
	"testing"
)

func TestNoConnectionToDatabase(t *testing.T) {
	t.Log("testing GetAllSongData before connection to database")
	GetAllSongData()
	t.Log("testing GetArtistSongData before connection to database")
	GetArtistSongData("Yorushika")
}

func TestConnectionToDatabase(t *testing.T) {
	t.Log("connecting to database")
	ConnectToDatabase()
	t.Log("testing GetAllSongData after connection to database")
	GetAllSongData()
	t.Log("testing GetArtistSongData after connection to database")
	GetArtistSongData("Yorushika")
	CloseDatabase()
}