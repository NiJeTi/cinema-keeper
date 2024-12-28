package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

type Config struct {
	ConnectionString string `conf:"connection_string"`
}

func Connect(cs string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
