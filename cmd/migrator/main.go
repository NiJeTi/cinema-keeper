package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/nijeti/cinema-keeper/internal/db"
	cfgPkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbUtils"
)

type config struct {
	DB db.Config `conf:"db"`
}

func main() {
	log.Println("starting")

	cfg := cfgPkg.ReadConfig[config]()

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")

	// db
	dbConn := db.Connect(cfg.DB)
	defer dbConn.Close()
	txWrapper := dbUtils.NewTxWrapper(dbLogger, dbConn)
	dbMigrator := db.NewMigrator(dbLogger, dbConn, txWrapper)

	// perform migrations
	log.Println("begin database migration")

	err := dbMigrator.Migrate()
	if err != nil {
		log.Fatalln("failed to migrate db:", err)
	}

	log.Println("database migration complete")
}
