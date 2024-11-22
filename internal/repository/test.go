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

//func (t *Test) Create(ctx context.Context, userId int64, input appmodels.TestOutput) (int64, error) {
//	const op = "internal.repository.test.Create()"
//	var id int64
//
//	stmt, err := t.db.PrepareContext(ctx, "INSERT INTO tests(user_id, theme) VALUES($1, $2) RETURNING id")
//	if err != nil {
//		return 0, fmt.Errorf("%s: %w", op, err)
//	}
//	defer stmt.Close()
//	res, err := stmt.ExecContext(ctx, stmt, userId, input.Theme)
//
//}

func (t *Test) Answer(ctx context.Context, userId int64, input appmodels.AnswersInput) (int64, error) {
	const op = "internal.repository.test.Answer()"
	query := "INSERT INTO tests (user_id, theme) VALUES ($1, $2) ON CONFLICT (user_id, theme) DO UPDATE SET score = score + $3 RETURNING id"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()
	res := stmt.QueryRowContext(ctx, userId, input.Theme, input.CorrectCount)
	var id int64
	if err := res.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

//func (t *Test) GetAllTests(ctx context.Context, userId int64) ([]appmodels.TestOutput, error) {
//	panic("implement me")
//}

//func (t *Test) GetTestById(ctx context.Context, userId int64, testId int64) (appmodels.TestOutput, error) {
//	panic("implement me")
//}
