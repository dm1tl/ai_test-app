package services

import (
	"ai_test-app/clients/testgen"
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/repository"
	"context"
	"fmt"
)

type TestService struct {
	repo repository.Test
	gen  testgen.TestGenerator
}

func NewTestService(repo repository.Test, gen testgen.TestGenerator) *TestService {
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
