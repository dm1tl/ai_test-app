package repository

import (
	appmodels "ai_test-app/internal/app_models"
	"context"
)

type Repository struct {
	Authorization
	Test
}

type Authorization interface {
	Create(ctx context.Context, userId int64, username string) error
}

type TestManager interface {
	Create(ctx context.Context, input appmodels.TestInput) error
	Answer(ctx context.Context, input appmodels.AnswersInput) error
	GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error)
	GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error)
}
