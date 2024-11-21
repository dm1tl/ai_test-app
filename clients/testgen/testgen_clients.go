package testgen

import (
	appmodels "ai_test-app/internal/app_models"
	"context"
)

type TestGenerator interface {
	Generate(ctx context.Context, input appmodels.TestInput) (appmodels.TestOutput, error)
}

type GenClient struct {
	client string
}

func (g *GenClient) Generate(ctx context.Context, input appmodels.TestInput) (appmodels.TestOutput, error) {
	return appmodels.TestOutput{}, nil
}
