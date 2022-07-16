package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("failed to load config.ini: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		LogFile: cfg.Section("scout_go").Key("log_file").String(),
	}
}