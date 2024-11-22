package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func executeQuery(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) error {
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
