package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schemaSQL string

// Open opens the database file, creating the schema if needed.
func Open(ctx context.Context, path string) (*sql.DB, *Queries, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, nil, fmt.Errorf("opening DB: %w", err)
	}

	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		return nil, nil, fmt.Errorf("creating DB schema: %w", err)
	}

	return db, New(db), nil
}
