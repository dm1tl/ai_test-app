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
	Create(ctx context.Context, userId int64) error
}

type TestManager interface {
	Create(ctx context.Context, input appmodels.TestInput) error
	Answer(ctx context.Context, input appmodels.AnswersInput) error
}
