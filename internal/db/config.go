package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

type Config struct {
	ConnectionString string `conf:"connection_string"`
}

func Connect(cs string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cs)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
