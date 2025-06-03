package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/NarthurN/TODO-API-web/internal/config"
	_ "modernc.org/sqlite"
)

type TaskStorage struct {
	SqlStorage *sql.DB
}

func New() (*TaskStorage, error) {
	dbFile := config.Cfg.TODO_DBFILE
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: cannot open database: %w", err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: failed to ping database: %w", err)
	}

	storage := &TaskStorage{SqlStorage: db}

	if err := createTable(storage); err != nil {
		return nil, fmt.Errorf("createTable: cannot create table: %w", err)
	}

	return storage, nil
}

func createTable(storage *TaskStorage) error {
	_, err := storage.SqlStorage.Exec(`
		CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date CHAR(8) NOT NULL DEFAULT "",
			title VARCHAR(256) NOT NULL DEFAULT "",
			comment TEXT NOT NULL DEFAULT "",
			repeat VARCHAR(128) NOT NULL DEFAULT ""
		);
		CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);
	`)
	if err != nil {
		return fmt.Errorf("storage.SqlStorage.Exec: failed to create scheduler table: %w", err)
	}

	return nil
}

func (t *TaskStorage) Close() error {
	err := t.SqlStorage.Close()
	if err != nil {
		return fmt.Errorf("t.SqlStorage.Close: error closing db: %w", err)
	}
	return nil
}
