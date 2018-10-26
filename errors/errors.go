package errors

import (
	"errors"
)

var (
	/*
		errors: config/config.go
	*/

	//ErrCfgNotFound is used when a configuration is not found and a default configuration will be used in its place
	ErrCfgNotFound = errors.New("configuration not found, using default config")
	//ErrCfgNoMOTD is used when a configuration lacks a MOTD
	ErrCfgNoMOTD = errors.New("no motd found")
	//ErrCfgNoSrvName is used when a configuration lacks a server name
	ErrCfgNoSrvName = errors.New("no serverName found")
	//ErrCfgInvalidSrvPort is used when a configuration either lacks a server port or has an invalid server port
	ErrCfgInvalidSrvPort = errors.New("invalid or no serverPort")
	//ErrCfgInvalidGameMode is used when a configuration either lacks a game mode or has an invalid game mode
	ErrCfgInvalidGameMode = errors.New("invalid or no gameMode")
	//ErrCfgInvalidMaxPlayers is used when a configuration either lacks a max players count or has an invalid max players count
	ErrCfgInvalidMaxPlayers = errors.New("invalid or no maxPlayers")
)
