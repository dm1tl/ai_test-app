package repository

import (
	appmodels "ai_test-app/internal/app_models"
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	TestManager
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TestManager:   NewTestRepository(db),
	}
}

type Authorization interface {
	Create(ctx context.Context, userId int64, username string) error
}

type TestManager interface {
	//	Create(ctx context.Context, userId int64, input appmodels.TestOutput) (int64, error)
	Answer(ctx context.Context, userId int64, input appmodels.AnswersInput) (int64, error)
	//	GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error)
	//	GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error)
}
