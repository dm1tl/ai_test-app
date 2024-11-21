package config

import (
	"time"

	"github.com/joho/godotenv"
)

type SSOConfig interface {
	Address() string
	Timeout() time.Duration
	RetriesCount() uint
}

type DBConfig interface {
	DSN() string
}

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
