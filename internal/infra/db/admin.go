package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateDatabaseIfNotExists(ctx context.Context, cfg Config) error {
	conn, err := pgx.Connect(ctx, cfg.AdminDSN())
	if err != nil {
		return fmt.Errorf("connect admin db: %w", err)
	}
	defer conn.Close(ctx)

	var exists bool
	if err := conn.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.Database).Scan(&exists); err != nil {
		return fmt.Errorf("check db exists: %w", err)
	}
	if exists {
		return nil
	}

	quoted := pgx.Identifier{cfg.Database}.Sanitize()
	if _, err := conn.Exec(ctx, "CREATE DATABASE "+quoted); err != nil {
		return fmt.Errorf("create database: %w", err)
	}
	return nil
}
