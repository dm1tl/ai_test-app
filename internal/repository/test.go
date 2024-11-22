package repository

import (
	appmodels "ai_test-app/internal/app_models"
	"context"
	"fmt"

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
	const op = "internal.repository.test.Create()"
	err := executeQuery(ctx, t.db, "INSERT INTO tests(id, user_id, theme) VALUES($1, $2, $3)", input.TestId, userId, input.Theme)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (t *Test) Answer(ctx context.Context, input appmodels.AnswersInput) error {
	const op = "internal.repository.test.Answer()"
	err := executeQuery(ctx, t.db, "UPDATE tests SET score = score + $1 WHERE id = $2 AND user_id = $3",
		input.CorrectCount, input.TestId, input.UserId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

//func (t *Test) GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error) {
//	panic("implement me")
//}

//func (t *Test) GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error) {
//	panic("implement me")
//}
