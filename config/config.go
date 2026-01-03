package config

import (
	"fmt"
	"path/filepath"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func LoadEnvFile() {
	// load environment variables (sql connection)
	envPath, err := findEnvFile()
	if err != nil {
		fmt.Println(".env file was not found")
		return
	}
	err = godotenv.Load(filepath.Join(envPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func findEnvFile() (string, error){
	envPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		_, err :=  os.Stat(filepath.Join(envPath, "go.mod"))
		if err == nil {
			return envPath, nil
		}

		parent := filepath.Dir(envPath)
		if parent == envPath {
			return "", fmt.Errorf("go.mod not found")
		}
		envPath = parent
	}
}