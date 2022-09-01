package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// LoadConfig loads configuration file for release-automator tools
func LoadConfig(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config %s: %w", path, err)
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshall config %s: %w", path, err)
	}

	// TODO: validate config

	return config, nil
}
