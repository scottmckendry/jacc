package config_test

import (
	"testing"

	"jacc/config"
)

const exampleConfigPath = "config.example.toml"

func TestLoadConfig(t *testing.T) {
	config, err := config.LoadConfig(exampleConfigPath)
	if err != nil {
		t.Errorf("Expected err to be nil, got '%s'", err)
	}
	if config.Name != "Jacc" {
		t.Errorf("Expected config.Name to be 'Jacc', got '%s'", config.Name)
	}
}

func TestLoadConfigInvalidPath(t *testing.T) {
	_, err := config.LoadConfig("invalid")
	if err == nil {
		t.Errorf("Expected err to be not nil, got nil")
	}
}

func TestLoadConfigInvalidFile(t *testing.T) {
	_, err := config.LoadConfig("config_test.go")
	if err == nil {
		t.Errorf("Expected err to be not nil, got nil")
	}
}
