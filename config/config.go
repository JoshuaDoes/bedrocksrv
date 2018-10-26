package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/JoshuaDoes/bedrocksrv/errors"
)

//Config stores a bedrocksrv configuration
type Config struct {
	MOTD            string `json:"motd"`            //Message of the Day
	SrvName         string `json:"serverName"`      //Name of the server
	SrvPort         int    `json:"serverPort"`      //Server connection port
	GameMode        int    `json:"gameMode"`        //Game mode (0 = survival, 1 = creative, 2 = adventure, 3 = spectator)
	MaxPlayers      int    `json:"maxPlayers"`      //Max amount of players at once
	SpawnProtection bool   `json:"spawnProtection"` //Whether to protect spawn from regular players
}

//Check checks whether a configuration has valid values
func (cfg *Config) Check() error {
	if cfg.MOTD == "" {
		return errors.ErrCfgNoMOTD
	}
	if cfg.SrvName == "" {
		return errors.ErrCfgNoSrvName
	}
	if cfg.SrvPort <= 0 || cfg.SrvPort > 65535 {
		return errors.ErrCfgInvalidSrvPort
	}
	if cfg.GameMode < 0 || cfg.GameMode > 3 {
		return errors.ErrCfgInvalidGameMode
	}
	if cfg.MaxPlayers <= 0 {
		return errors.ErrCfgInvalidMaxPlayers
	}

	return nil
}

//Load returns the configuration in the requested file or an error otherwise
func (cfg *Config) Load(path string) error {
	cfgFile, err := os.Open(path)
	if err != nil {
		err = defaultConfig.Save(path)
		if err != nil {
			return err
		}
		*cfg = *defaultConfig
		return errors.ErrCfgNotFound
	}
	defer cfgFile.Close()

	cfgDecoder := json.NewDecoder(cfgFile)

	err = cfgDecoder.Decode(&cfg)
	if err != nil {
		return err
	}

	err = cfg.Check()
	if err != nil {
		return err
	}

	return nil
}

//Save saves the configuration to the requested file or an error otherwise
func (cfg *Config) Save(path string) error {
	cfgJSON, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, cfgJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

var defaultConfig = &Config{
	MOTD:            "A Minecraft server by Gophers",
	SrvName:         "bedrocksrv",
	SrvPort:         19132,
	GameMode:        0,
	MaxPlayers:      20,
	SpawnProtection: true,
}
