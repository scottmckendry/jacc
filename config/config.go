package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Theme   string
	Profile struct {
		Name        string
		TagLine     string
		Description string
		Skills      []string
	}
	Contact struct {
		Enabled      bool
		Email        string
		Phone        string
		Municipality string
	}
	Links struct {
		Enabled   bool
		GitHub    string
		LinkedIn  string
		Twitter   string
		Instagram string
		Facebook  string
		Website   string
	}
	Education []struct {
		Enabled    bool
		Name       string
		Credential string
		StartDate  string
		EndDate    string
	}
	Experience []struct {
		Enabled      bool
		Name         string
		Title        string
		Description  string
		StartDate    string
		EndDate      string
		Achievements []string
	}
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
