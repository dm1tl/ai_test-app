package services

import (
	appmodels "ai_test-app/internal/app_models"
	"context"
)

type Service struct {
	Authorization
	Test
}

type Authorization interface {
	Create(ctx context.Context, user appmodels.User) error
	Login(ctx context.Context, input appmodels.SignInInput) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
}

type Test interface {
	Create(ctx context.Context, input appmodels.TestInput) (appmodels.TestOutput, error)
	Answer(ctx context.Context, input appmodels.AnswersInput) error
}
