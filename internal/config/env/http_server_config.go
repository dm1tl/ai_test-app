package env

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type HTTPServer struct {
	Address        string        `yaml:"address"`
	MaxHeaderBytes int           `yaml:"maxheaderbytes"`
	ReadTimeout    time.Duration `yaml:"readtimeout"`
	WriteTimeout   time.Duration `yaml:"writetimeout"`
	IdleTimeout    time.Duration `yaml:"idletimeout"`
}

func NewHTTPServerConfig() (*HTTPServer, error) {
	var config HTTPServer
	filepath := os.Getenv("HTTPSERVER_CONFIG_PATH")
	if filepath == "" {
		return nil, fmt.Errorf("HTTPSERVER_CONFIG_PATH is empty")
	}
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load config file #%v", err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("couldn't parse config file into model #%v", err)
	}
	return &config, nil
}
