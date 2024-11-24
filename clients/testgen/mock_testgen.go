package testgen

import (
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/config/env"
	"context"
)

type MockTestGenServiceClient struct{}

func NewMockTestGenServiceClient(cfg *env.GenTestConfig) *MockTestGenServiceClient {
	return &MockTestGenServiceClient{}
}

func (m *MockTestGenServiceClient) Generate(ctx context.Context, input appmodels.TestInput) (appmodels.TestOutput, error) {
	mockAnswer := appmodels.Answer{
		AnswerId:  1,
		AnswerTxt: "mock-answer",
		IsCorrect: true,
	}
	mockQuestion := appmodels.Question{
		QuestionId: 1,
		Question:   "mock-question",
		Answers:    []appmodels.Answer{mockAnswer},
	}
	output := appmodels.TestOutput{
		TestId:    1,
		Theme:     "mock-theme",
		Questions: []appmodels.Question{mockQuestion},
	}
	return output, nil
}
