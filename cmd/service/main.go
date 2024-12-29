package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/nijeti/healthcheck"
	"github.com/nijeti/healthcheck/servers/fasthttp"

	"github.com/nijeti/cinema-keeper/internal/db"
	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/quote"
	"github.com/nijeti/cinema-keeper/internal/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
	cfgpkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbutils"
)

const (
	codeOk  = 0
	codeErr = 1
)

type config struct {
	Discord discord.Config `conf:"discord"`
	DB      db.Config      `conf:"db"`
}

func main() {
	code := run()
	os.Exit(code)
}

func run() int {
	log.Println("starting")
	defer log.Println("shutdown complete")

	cfg, err := cfgpkg.ReadConfig[config]()
	if err != nil {
		log.Println("failed to read config:", err)
		return codeErr
	}

	// todo: https://github.com/NiJeTi/cinema-keeper/issues/29
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")
	cmdLogger := logger.WithGroup("command")
	hcLogger := logger.WithGroup("healthcheck")

	// db
	dbConn, err := db.Connect(cfg.DB.ConnectionString)
	if err != nil {
		log.Println("failed to connect to db:", err)
		cancel()
		return codeErr
	}
	defer dbConn.Close()
	dbProbe := db.NewProbe(dbConn.DB)
	txWrapper := dbutils.NewTxWrapper(dbLogger, dbConn.DB)

	quotesRepo := db.NewQuotesRepo(dbLogger, dbConn.DB, txWrapper)

	// discord
	discordSession, err := discord.Connect(cfg.Discord.Token)
	if err != nil {
		log.Println("failed to connect to Discord:", err)
		cancel()
		return codeErr
	}
	defer discordSession.Close()

	commands := map[string]*discord.Command{
		discord.QuoteName: {
			Description: discord.Quote,
			Handler:     quote.New(ctx, cmdLogger, discordSession, quotesRepo),
		},
		discord.CastName: {
			Description: discord.Cast,
			Handler:     cast.New(ctx, cmdLogger, discordSession),
		},
		discord.LockName: {
			Description: discord.Lock,
			Handler:     lock.New(ctx, cmdLogger, discordSession),
		},
		discord.UnlockName: {
			Description: discord.Unlock,
			Handler:     unlock.New(ctx, cmdLogger, discordSession),
		},
		discord.RollName: {
			Description: discord.Roll,
			Handler:     roll.New(ctx, cmdLogger, discordSession),
		},
	}

	err = discord.RegisterCommands(discordSession, commands, cfg.Discord.Guild)
	if err != nil {
		log.Println("failed to register commands:", err)
		cancel()
		return codeErr
	}
	//nolint:errcheck // call is deferred
	defer discord.UnregisterCommands(
		discordSession, commands, cfg.Discord.Guild,
	)

	discordProbe := discord.NewProbe(discordSession)

	// healthcheck
	hc := healthcheck.New(
		healthcheck.WithLogger(hcLogger),
		healthcheck.WithProbe("db", dbProbe),
		healthcheck.WithProbe("discord", discordProbe),
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
	log.Println("startup complete")
	<-stop

	// shutdown
	log.Println("shutting down")
	cancel()

	return codeOk
}
