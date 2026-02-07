package db

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func MigrateUp(ctx context.Context, cfg Config, migrationsDir string) error {
	db, err := sql.Open("pgx", cfg.AppDSN())
	if err != nil {
		return fmt.Errorf("open db for migration: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping db for migration: %w", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}
	if err := goose.UpContext(ctx, db, filepath.Clean(migrationsDir)); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}
	return nil
}
