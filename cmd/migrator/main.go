package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pressly/goose/v3"

	"github.com/nijeti/cinema-keeper/internal/db"
	cfgPkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
)

type config struct {
	DB db.Config `conf:"db"`
}

const (
	codeOk  = 0
	codeErr = 1
)

const migrationsTable = "migrations"

//go:embed *.sql
var migrations embed.FS

func main() {
	code := run()
	os.Exit(code)
}

func run() (code int) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic", "error", err)
			code = codeErr
		}
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger.Info("starting")

	// config
	cfg, err := cfgPkg.ReadConfig[config]()
	if err != nil {
		logger.ErrorContext(ctx, "failed to read config", "error", err)
		return codeErr
	}

	// db
	dbConn, err := db.Connect(ctx, cfg.DB)
	if err != nil {
		logger.ErrorContext(ctx, "failed to connect to db", "error", err)
		return codeErr
	}
	defer dbConn.Close()

	// migrator
	goose.SetLogger(gooseLogger{logger})
	goose.SetBaseFS(migrations)
	goose.SetTableName(migrationsTable)
	goose.SetSequential(true)

	err = goose.SetDialect(string(goose.DialectPostgres))
	if err != nil {
		logger.ErrorContext(ctx, "failed to set dialect", "error", err)
		return codeErr
	}

	// run
	logger.Info("begin database migration")

	err = goose.UpContext(ctx, dbConn.DB, ".")
	if err != nil {
		logger.ErrorContext(ctx, "failed to migrate database", "error", err)
		return codeErr
	}

	logger.Info("database migration complete")
	return codeOk
}
