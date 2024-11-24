package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// executeQuery executes an SQL query within a transaction, checks the result of execution and returns an error if something went wrong.
func executeQueryTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) error {
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return err
	}
	check, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if check != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", check)
	}
	return nil
}

// executeQuery executes an SQL query and returns an error if something went wrong.
func executeQuery(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	check, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if check != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", check)
	}
	return nil
}
