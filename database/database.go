package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
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

func GetAllSongData() ([]SongData, error) {
	if !checkDatabaseConnection() {
		return nil, fmt.Errorf("Not connected to a database")
	}
	results, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err.Error())
	}
	
	defer results.Close()

	var songsSlice []SongData

	for results.Next() {
		var song SongData
		err := results.Scan(&song.ID, &song.Artist, &song.Song_name, &song.Song_file_name, &song.Duration)
		if err != nil {
			panic(err.Error())
		}
		songsSlice = append(songsSlice, song)
	}
	if err = results.Err(); err != nil {
		panic(err.Error())
	}

	return songsSlice, nil
}

func GetArtistSongData(artistName string) ([]SongData, error) {
	if !checkDatabaseConnection() {
		return nil, fmt.Errorf("Not connected to a database")
	}
	results, err := db.Query("SELECT * FROM songs where artist = ?", artistName)
	if err != nil {
		panic(err.Error())
	}
	
	defer results.Close()

	var songsSlice []SongData

	for results.Next() {
		var song SongData
		err := results.Scan(&song.ID, &song.Artist, &song.Song_name, &song.Song_file_name, &song.Duration)
		if err != nil {
			panic(err.Error())
		}
		songsSlice = append(songsSlice, song)
	}
	if err = results.Err(); err != nil {
		panic(err.Error())
	}

	return songsSlice, nil
}

func checkDatabaseConnection() bool {
	if db == nil {
		fmt.Println("Not connected to a database, returning")
		return false
	} else {
		return true
	}
}

