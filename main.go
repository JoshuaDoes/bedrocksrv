package main

import (
	"io/ioutil"
	"os"

	"github.com/JoshuaDoes/bedrocksrv/config"
	"github.com/JoshuaDoes/bedrocksrv/errors"
)

var (
	cfg = &config.Config{}
)

func main() {
	initLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("bedrocksrv © JoshuaDoes: 2018.")
	Info.Println(BuildID)

	Info.Println("> Loading server configuration...")
	err := cfg.Load("config.json")
	if err != nil {
		if err == errors.ErrCfgNotFound {
			Warning.Println("Configuration not found, using default configuration")
		} else {
			Error.Println(err)
		}
	}

	Info.Println("-- [START] DISPLAY CONFIG --")
	Info.Printf("Message of the Day: %s\n", cfg.MOTD)
	Info.Printf("Server Name: %s\n", cfg.SrvName)
	Info.Printf("Server Port: %d\n", cfg.SrvPort)
	Info.Printf("Game Mode: %d\n", cfg.GameMode)
	Info.Printf("Max Players: %d\n", cfg.MaxPlayers)
	Info.Printf("Spawn Protection: %v\n", cfg.SpawnProtection)
	Info.Println("-- [END] DISPLAY CONFIG --")
}
