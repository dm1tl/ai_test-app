package env

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ssoConfig struct {
	address      string        `yaml:"address"`
	timeout      time.Duration `yaml:"timeout"`
	retriesCount uint          `yaml:"retriescount"`
}

func (s *ssoConfig) Address() string {
	return s.address
}

func (s *ssoConfig) Timeout() time.Duration {
	return s.timeout
}

func (s *ssoConfig) RetriesCount() uint {
	return s.retriesCount
}

func NewSSOConfig() (*ssoConfig, error) {
	var config ssoConfig
	filepath := os.Getenv("SSO_CONFIG_PATH")
	if filepath == "" {
		return nil, fmt.Errorf("SSO_CONFIG_PATH is empty")
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
