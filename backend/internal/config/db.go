package config

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(databaseURL string) (*sql.DB, error) {
	// Parse sqlite://./data/bukadulu.db → ./data/bukadulu.db
	path := strings.TrimPrefix(databaseURL, "sqlite://")

	// Ensure directory exists
	dir := path
	if idx := strings.LastIndex(dir, "/"); idx > 0 {
		dir = dir[:idx]
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("create db dir: %w", err)
		}
	}

	db, err := sql.Open("sqlite3", path+"?_journal_mode=WAL&_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	// Enable WAL mode for concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		slog.Warn("failed to set WAL mode", "error", err)
	}
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		slog.Warn("failed to enable foreign keys", "error", err)
	}

	slog.Info("database connected", "path", path)
	return db, nil
}

func RunMigrations(db *sql.DB, migrationPath string) error {
	sqlBytes, err := os.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	statements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("migration failed on: %s... : %w", stmt[:min(len(stmt), 50)], err)
		}
	}

	slog.Info("migrations applied successfully")
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
