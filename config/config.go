package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type ApiConfig struct {
	Token, Host, Cert, Key string
	ChatId int64
}

func ReadConfig() ApiConfig {
	var conf ApiConfig
    if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
        log.Fatal(err)
    }
	return conf
}