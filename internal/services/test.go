package services

import (
	"ai_test-app/clients/testgen"
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/repository"
	"context"
	"fmt"
)

type TestService struct {
	repo repository.TestManager
	gen  testgen.TestGenerator
}

func NewTestService(repo repository.TestManager, gen testgen.TestGenerator) *TestService {
	return &TestService{
		repo: repo,
		gen:  gen,
	}
}

func (t *TestService) Create(ctx context.Context, input appmodels.TestInput) (*appmodels.TestOutput, error) {
	const op = "internal.services.Create()"
	var output appmodels.TestOutput
	test, err := t.gen.Generate(ctx, input)
	if err != nil {
		return &output, fmt.Errorf("%s: %w", op, err)
	}
	if err := t.repo.Create(ctx, input); err != nil {
		return nil, nil
	}
	return &test, nil
}

func (t *TestService) Answer(ctx context.Context, input appmodels.AnswersInput) error {
	const op = "internal.services.Answer()"
	err := t.repo.Answer(ctx, input)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (t *TestService) GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error) {
	const op = "internal.services.GetAllTests()"
	output, err := t.repo.GetAllTests(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return output, nil
}

func (t *TestService) GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error) {
	const op = "internal.services.GetTestById()"
	var output appmodels.TestOutput
	output, err := t.repo.GetTestById(ctx, userId, testId)
	if err != nil {
		return output, fmt.Errorf("%s: %w", op, err)
	}
	return output, nil
}
