package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/nijeti/healthcheck"
	hcHTTP "github.com/nijeti/healthcheck/servers/http"

	dcAdapterPkg "github.com/nijeti/cinema-keeper/internal/adapters/discord"
	"github.com/nijeti/cinema-keeper/internal/adapters/omdb"
	"github.com/nijeti/cinema-keeper/internal/db"
	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/movie"
	"github.com/nijeti/cinema-keeper/internal/handlers/quote"
	"github.com/nijeti/cinema-keeper/internal/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
	cfgPkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/services/addMovie"
	"github.com/nijeti/cinema-keeper/internal/services/addQuote"
	"github.com/nijeti/cinema-keeper/internal/services/diceRoll"
	"github.com/nijeti/cinema-keeper/internal/services/listUserQuotes"
	"github.com/nijeti/cinema-keeper/internal/services/lockVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/services/mentionVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/services/presence"
	"github.com/nijeti/cinema-keeper/internal/services/printRandomQuote"
	"github.com/nijeti/cinema-keeper/internal/services/searchNewMovie"
	"github.com/nijeti/cinema-keeper/internal/services/unlockVoiceChan"
)

type config struct {
	Discord discord.Config `conf:"discord"`
	DB      db.Config      `conf:"db"`
	OMDB    omdb.Config    `conf:"omdb"`
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

	logger.Info("starting")

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

	dbProbe := db.NewProbe(dbConn.DB)

	quotesRepo := db.NewQuotesRepo(dbConn)
	moviesRepo := db.NewMoviesRepo(dbConn)

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

	dcAdapter := dcAdapterPkg.New(dcRouter.Session())
	omdbAdapter := omdb.New(cfg.OMDB)

	addMovieSvc := addMovie.New(dcAdapter, omdbAdapter, moviesRepo)
	addQuoteSvc := addQuote.New(dcAdapter, quotesRepo)
	diceRollSvc := diceRoll.New(dcAdapter)
	listUserQuotesSvc := listUserQuotes.New(dcAdapter, quotesRepo)
	lockVoiceChanSvc := lockVoiceChan.New(dcAdapter)
	mentionVoiceChanSvc := mentionVoiceChan.New(dcAdapter)
	presenceSvc := presence.New(dcAdapter)
	printRandomQuoteSvc := printRandomQuote.New(quotesRepo, dcAdapter)
	searchNewMovieSvc := searchNewMovie.New(dcAdapter, omdbAdapter)
	unlockVoiceChanSvc := unlockVoiceChan.New(dcAdapter)

	err = dcRouter.SetCommands(
		discord.Command{
			Description: commands.Quote(),
			Handler: quote.New(
				printRandomQuoteSvc, listUserQuotesSvc, addQuoteSvc,
			),
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
			Handler:     roll.New(diceRollSvc),
		},
		discord.Command{
			Description: commands.Movie(),
			Handler:     movie.New(searchNewMovieSvc, addMovieSvc),
		},
	)
	if err != nil {
		logger.ErrorContext(ctx, "failed to set commands", "error", err)
		return codeErr
	}

	dcProbe := discord.NewProbe(dcRouter.Session())

	// healthcheck
	hcLogger := logger.WithGroup("healthcheck")

	hc := healthcheck.New(
		healthcheck.WithLogger(hcLogger),
		healthcheck.WithProbe("db", dbProbe),
		healthcheck.WithProbe("discord", dcProbe),
	)
	hcs := hcHTTP.New(
		hc,
		hcHTTP.WithLogger(hcLogger),
		hcHTTP.WithAddress(":8080"),
		hcHTTP.WithRoute("/health"),
	)
	hcs.Start()
	defer hcs.Stop()

	// run
	logger.Info("startup complete")

	err = presenceSvc.Set(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "failed to set presence", "error", err)
	}

	<-ctx.Done()

	// shutdown
	logger.Info("shutting down")

	logger.Info("shutdown complete")
	return codeOk
}
