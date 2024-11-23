package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

//func (a *Auth) Create(ctx context.Context, userId int64, username string) error {
//	const op = "internal.repository.Create()"
//	tx, err := a.db.BeginTx(ctx, nil)
//	if err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//	defer func() {
//		if err != nil {
//			_ = tx.Rollback()
//		}
//	}()
//	if err := executeQuery(ctx, tx, "INSERT INTO users (id, username) VALUES($1, $2)", userId, username); err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//
//	if err := executeQuery(ctx, tx, "INSERT INTO tests (user_id) VALUES($1)", userId); err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//
//	if err := tx.Commit(); err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//	return nil
//}

func (a *AuthRepository) Create(ctx context.Context, userId int64, username string) error {
	const op = "internal.repository.Create()"
	err := executeQuery(ctx, a.db, "INSERT INTO users(id, username) VALUES ($1, $2)", userId, username)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
