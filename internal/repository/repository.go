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
	Login(ctx context.Context, input appmodels.SignInInput) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
}

type Test interface {
	Create(ctx context.Context, input appmodels.TestInput) error
	Answer(ctx context.Context, input appmodels.AnswersInput) error
}
