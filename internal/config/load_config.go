package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"os"
)

// LoadConfig loads configuration file for release-automator tools
func LoadConfig(path string) (Config, error) {
	content, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config %s: %w", path, err)
	}

	config := Config{}

	yamlDecoder := yaml.NewDecoder(content)
	yamlDecoder.KnownFields(true)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshall config %s: %w", path, err)
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to validate config %s: %w", path, err)
	}

	return config, nil
}
