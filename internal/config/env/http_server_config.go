package env

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		Handler:        handler,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		IdleTimeout:    cfg.IdleTimeout,
	}
	return s.httpServer.ListenAndServe()
}

type HTTPServer struct {
	Address        string        `yaml:"address"`
	MaxHeaderBytes int           `yaml:"maxheaderbytes"`
	ReadTimeout    time.Duration `yaml:"readtimeout"`
	WriteTimeout   time.Duration `yaml:"writetimeout"`
	IdleTimeout    time.Duration `yaml:"idletimeout"`
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewHTTPServerConfig() (*HTTPServer, error) {
	var config HTTPServer
	filepath := os.Getenv("HTTP_SERVER_CONFIG_PATH")
	if filepath == "" {
		return nil, fmt.Errorf("HTTP_SERVER_CONFIG_PATH is empty")
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
