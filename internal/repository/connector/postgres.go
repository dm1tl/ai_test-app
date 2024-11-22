package connector

import (
	"ai_test-app/internal/config"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
