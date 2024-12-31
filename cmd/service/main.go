package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/nijeti/healthcheck"
	"github.com/nijeti/healthcheck/servers/fasthttp"

	dcadapter "github.com/nijeti/cinema-keeper/internal/adapters/discord"
	"github.com/nijeti/cinema-keeper/internal/db"
	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/quote"
	"github.com/nijeti/cinema-keeper/internal/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
	cfgpkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/services/addQuote"
	"github.com/nijeti/cinema-keeper/internal/services/diceRoll"
	"github.com/nijeti/cinema-keeper/internal/services/listQuotes"
	"github.com/nijeti/cinema-keeper/internal/services/lockVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/services/mentionVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/services/unlockVoiceChan"
)

type config struct {
	Discord discord.Config `conf:"discord"`
	DB      db.Config      `conf:"db"`
}

const (
	codeOk  = 0
	codeErr = 1
)

func main() {
	code := run()
	os.Exit(code)
}

func run() (code int) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic", "error", err)
		}

		code = codeErr
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger.Info("starting")
	defer logger.Info("shutdown complete")

	cfg, err := cfgpkg.ReadConfig[config]()
	if err != nil {
		logger.ErrorContext(ctx, "failed to read config", "error", err)
		return codeErr
	}

	// db
	dbLogger := logger.WithGroup("db")

	dbConn, err := db.Connect(ctx, cfg.DB)
	if err != nil {
		logger.ErrorContext(ctx, "failed to connect to db", "error", err)
		return codeErr
	}

	dbProbe := db.NewProbe(dbConn.DB)

	quotesRepo := db.NewQuotesRepo(dbLogger, dbConn)

	// discord
	dcLogger := logger.WithGroup("discord")

	dcRouter, err := discord.NewRouter(
		cfg.Discord,
		discord.WithContext(ctx),
		discord.WithLogger(dcLogger),
	)
	if err != nil {
		logger.ErrorContext(
			ctx, "failed to create Discord router", "error", err,
		)
		return codeErr
	}
	defer dcRouter.Close()

	dcAdapter := dcadapter.New(dcRouter.Session())

	addQuoteSvc := addQuote.New(dcAdapter, quotesRepo)
	listQuotesSvc := listQuotes.New(dcAdapter, quotesRepo)
	lockVoiceChanSvc := lockVoiceChan.New(dcAdapter)
	mentionVoiceChanSvc := mentionVoiceChan.New(dcAdapter)
	rollSvc := diceRoll.New(dcAdapter)
	unlockVoiceChanSvc := unlockVoiceChan.New(dcAdapter)

	err = dcRouter.SetCommands(
		discord.Command{
			Description: commands.Quote(),
			Handler:     quote.New(listQuotesSvc, addQuoteSvc),
		},
		discord.Command{
			Description: commands.Cast(),
			Handler:     cast.New(mentionVoiceChanSvc),
		},
		discord.Command{
			Description: commands.Lock(),
			Handler:     lock.New(lockVoiceChanSvc),
		},
		discord.Command{
			Description: commands.Unlock(),
			Handler:     unlock.New(unlockVoiceChanSvc),
		},
		discord.Command{
			Description: commands.Roll(),
			Handler:     roll.New(rollSvc),
		},
	)
	if err != nil {
		logger.ErrorContext(ctx, "failed to set commands", "error", err)
		return codeErr
	}
	defer dcRouter.UnsetCommands() //nolint:errcheck // shutdown

	dcProbe := discord.NewProbe(dcRouter.Session())

	// healthcheck
	hcLogger := logger.WithGroup("healthcheck")

	hc := healthcheck.New(
		healthcheck.WithLogger(hcLogger),
		healthcheck.WithProbe("db", dbProbe),
		healthcheck.WithProbe("discord", dcProbe),
	)
	hcs := fasthttp.New(
		hc,
		fasthttp.WithLogger(hcLogger),
		fasthttp.WithAddress(":8080"),
		fasthttp.WithRoute("/health"),
	)
	hcs.Start()
	defer hcs.Stop()

	// run
	logger.Info("startup complete")
	<-ctx.Done()

	// shutdown
	logger.Info("shutting down")

	return codeOk
}
