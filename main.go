package main

import (
	"github.com/stanleyyellowzx/MusicGuesser/config"
	"github.com/stanleyyellowzx/MusicGuesser/client"
)

func main() {
	config.LoadEnvFile()
	client.Start()
}