package migration

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed */*.sql
var embedMigrations embed.FS

func MigrateTables(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Up(db, "user"); err != nil {
		return fmt.Errorf("up migration: %w", err)
	}

	if err := goose.Up(db, "kemono"); err != nil {
		return fmt.Errorf("up migration: %w", err)
	}

	return nil
}

func ResetKemonoTable(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Down(db, "kemono"); err != nil {
		return fmt.Errorf("down migration: %w", err)
	}

	if err := goose.Up(db, "kemono"); err != nil {
		return fmt.Errorf("up migration: %w", err)
	}

	return nil
}
