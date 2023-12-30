package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Name string
}

// LoadConfig loads the config from a given TOML file
func LoadConfig(filePath string) (Config, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
