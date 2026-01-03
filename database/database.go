package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"log"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dataSourceName := os.Getenv("DATABASE_CONNECTION")

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr.Error())
	}

	fmt.Println("Successfully Connected to MySQL database")

}

func CloseDatabase() {
	if db == nil {
		fmt.Println("Not connected to a database, returning")
	}

	fmt.Println("Closing database")
	db.Close()
}

func GetAllSongData() {
	if !checkDatabaseConnection() {
		fmt.Println("Not connected to a database, returning")
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
		fmt.Println("Not connected to a database, returning")
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
		return false
	} else {
		return true
	}
}

