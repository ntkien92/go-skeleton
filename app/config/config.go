package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"database" validate:"required"`
}

func NewConfig(filePath string) (Config, error) {
	var (
		config = Config{}
		err    error
	)

	if filePath == "" {
		return Config{}, fmt.Errorf("file path is empty")
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read YAML file: %w", err)
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return config, err
}
