package env

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type GenTestConfig struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount uint          `yaml:"retriescount"`
}

func NewGenTestConfig() (*GenTestConfig, error) {
	var config GenTestConfig
	filepath := os.Getenv("GEN_CONFIG_PATH")
	if filepath == "" {
		return nil, fmt.Errorf("GEN_CONFIG_PATH is empty")
	}
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load config file %w", err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("couldn't parse config file into model %w", err)
	}
	return &config, nil
}
