package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect(cs string) *sql.DB {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		log.Fatalln("failed to open database connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("failed to ping database", err)
	}

	return db
}
