package repository

import (
	appmodels "ai_test-app/internal/app_models"
	"context"

	"github.com/jmoiron/sqlx"
)

type Test struct {
	db *sqlx.DB
}

func NewTest(db *sqlx.DB) *Test {
	return &Test{
		db: db,
	}
}

func (t *Test) Create(ctx context.Context, userId int64, input appmodels.TestOutput) error {
	panic("implement me")
}

func (t *Test) Answer(ctx context.Context, input appmodels.AnswersInput) error {
	panic("implement me")
}

func (t *Test) GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error) {
	panic("implement me")
}

func (t *Test) GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error) {
	panic("implement me")
}
