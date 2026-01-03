package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/stanleyyellowzx/MusicGuesser/config"
)

type SongData struct {
	ID int
	Artist string
	Song_name string
	Song_file_name string
	Duration int
}

var db *sql.DB

func ConnectToDatabase() {
	// load environment variables (sql connection)
	config.LoadEnvFile()
	dataSourceName := os.Getenv("DATABASE_CONNECTION")
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully Connected to MySQL database")

}

func CloseDatabase() {
	if !checkDatabaseConnection() {
		return
	}

	fmt.Println("Closing database")
	db.Close()
}

func GetAllSongData() {
	if !checkDatabaseConnection() {
		return
	}
	results, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err.Error())
	}
	
	defer results.Close()

	var songsArray []SongData

	for results.Next() {
		var song SongData
		err := results.Scan(&song.ID, &song.Artist, &song.Song_name, &song.Song_file_name, &song.Duration)
		if err != nil {
			panic(err.Error())
		}
		songsArray = append(songsArray, song)
	}
	if err = results.Err(); err != nil {
		panic(err.Error())
	}

	for index, element := range songsArray {
		fmt.Println("Index: ", index, ", Element: ", element)
	}
}

func GetArtistSongData(artistName string) {
	if !checkDatabaseConnection() {
		return
	}
	results, err := db.Query("SELECT * FROM songs where artist = ?", artistName)
	if err != nil {
		panic(err.Error())
	}
	
	defer results.Close()

	var songsArray []SongData

	for results.Next() {
		var song SongData
		err := results.Scan(&song.ID, &song.Artist, &song.Song_name, &song.Song_file_name, &song.Duration)
		if err != nil {
			panic(err.Error())
		}
		songsArray = append(songsArray, song)
	}
	if err = results.Err(); err != nil {
		panic(err.Error())
	}

	for index, element := range songsArray {
		fmt.Println("Index: ", index, ", Element: ", element)
	}
}

func checkDatabaseConnection() bool {
	if db == nil {
		fmt.Println("Not connected to a database, returning")
		return false
	} else {
		return true
	}
}

