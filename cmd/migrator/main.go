package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/nijeti/cinema-keeper/internal/db"
	cfgpkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbutils"
)

const (
	codeOk  = 0
	codeErr = 1
)

type config struct {
	DB db.Config `conf:"db"`
}

func main() {
	code := run()
	os.Exit(code)
}

func run() int {
	log.Println("starting")

	cfg, err := cfgpkg.ReadConfig[config]()
	if err != nil {
		log.Println("failed to read config:", err)
		return codeErr
	}

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")

	// db
	dbConn, err := db.Connect(cfg.DB.ConnectionString)
	if err != nil {
		log.Println("failed to connect to db:", err)
		return codeErr
	}
	defer dbConn.Close()
	txWrapper := dbutils.NewTxWrapper(dbLogger, dbConn)
	dbMigrator := db.NewMigrator(dbLogger, dbConn, txWrapper)

	// perform migrations
	log.Println("begin database migration")

	err = dbMigrator.Migrate()
	if err != nil {
		log.Println("failed to migrate db:", err)
		return codeErr
	}

	log.Println("database migration complete")
	return codeOk
}
