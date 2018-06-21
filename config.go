package main

import (
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	DB Database `toml:"database"`
}

func InitConfig() {
    _, err := toml.DecodeFile( "./config.toml", &config)
    CheckError(err)
}
