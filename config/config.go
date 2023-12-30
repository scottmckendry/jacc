package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Name string
}

// LoadConfig loads the config from config.toml
func LoadConfig() Config {
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
