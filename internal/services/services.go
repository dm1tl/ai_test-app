package services

import (
	"ai_test-app/clients/sso"
	"ai_test-app/clients/testgen"
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/repository"
	"context"
)

type Service struct {
	Authorization
	Test
}

func NewService(repos *repository.Repository, ssoclient *sso.SSOClient, testm *testgen.GenClient) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, ssoclient.SSOProvider),
		Test:          NewTestService(repos.TestManager, testm.TestGenerator),
	}
}

type Authorization interface {
	Create(ctx context.Context, user appmodels.User) error
	Login(ctx context.Context, input appmodels.SignInInput) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
}

type Test interface {
	Create(ctx context.Context, userId int64, input appmodels.TestInput) (appmodels.TestOutput, error)
	Answer(ctx context.Context, userId int64, input appmodels.AnswersInput) (int64, error)
	//	GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error)
	//	GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error)
}
