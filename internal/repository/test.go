package repository

import "github.com/jmoiron/sqlx"

type Test struct {
	db *sqlx.DB
}

func NewTest(db *sqlx.DB) *Test {
	return &Test{
		db: db,
	}
}
