package tools

import (
	"context"
	"database/sql"
)

func CheckIfTableExists(db *sql.DB, ctx context.Context, tableCheck string) (bool, error) {
	rows, err := db.QueryContext(ctx, tableCheck)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return rows.Next(), nil
}

func CreateTable(db *sql.DB, ctx context.Context, ddl string) error {
	_, err := db.ExecContext(ctx, ddl)
	return err
}
