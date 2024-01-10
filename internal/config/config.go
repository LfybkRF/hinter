package config

import (
	"fmt"
	"l0_ms/internal/app/server"

	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database      DatabaseConfig      `yaml:"database"`
	Http          server.Config         `yaml:"http"`
	NatsStreaming NatsStreamingConfig `yaml:"nats_streaming"`
}

func NewConfig(configFile string) (*Config, error) {
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading file error: %w", err)
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(rawYAML, cfg); err != nil {
		return nil, fmt.Errorf("yaml parsing error: %w", err)
	}

	return cfg, nil
}