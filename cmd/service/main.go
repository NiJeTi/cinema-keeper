package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/nijeti/cinema-keeper/internal/db"
	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/quote"
	"github.com/nijeti/cinema-keeper/internal/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
	cfgPkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbUtils"
	"github.com/nijeti/cinema-keeper/internal/pkg/healthcheck"
)

type config struct {
	Discord discord.Config `conf:"discord"`
	DB      db.Config      `conf:"db"`
}

func main() {
	log.Println("starting")
	defer log.Println("shutdown complete")

	cfg := cfgPkg.ReadConfig[config]()

	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")
	cmdLogger := logger.WithGroup("command")
	hcLogger := logger.WithGroup("healthcheck")

	// db
	dbConn := db.Connect(cfg.DB)
	defer dbConn.Close()
	dbProbe := db.NewProbe(dbConn)
	txWrapper := dbUtils.NewTxWrapper(dbLogger, dbConn)

	// discord
	discordConn := discord.Connect(cfg.Discord.Token)
	defer discordConn.Close()
	discordProbe := discord.NewProbe(discordConn)

	// repos
	quotesRepo := db.NewQuotesRepo(dbLogger, dbConn, txWrapper)

	// handlers
	cmds := map[string]*discord.Command{
		discord.QuoteName: {
			Description: discord.Quote,
			Handler:     quote.New(ctx, cmdLogger, quotesRepo),
		},
		discord.CastName: {
			Description: discord.Cast,
			Handler:     cast.New(ctx, cmdLogger),
		},
		discord.LockName: {
			Description: discord.Lock,
			Handler:     lock.New(ctx, cmdLogger),
		},
		discord.UnlockName: {
			Description: discord.Unlock,
			Handler:     unlock.New(ctx, cmdLogger),
		},
		discord.RollName: {
			Description: discord.Roll,
			Handler:     roll.New(ctx, cmdLogger),
		},
	}

	// handlers
	discord.RegisterCommands(discordConn, cmds, cfg.Discord.Guild)
	defer discord.UnregisterCommands(discordConn, cmds, cfg.Discord.Guild)

	// healthcheck
	hc := healthcheck.New(
		hcLogger,
		healthcheck.WithProbe("db", dbProbe),
		healthcheck.WithProbe("discord", discordProbe),
	)
	hcServer := hc.Serve(":8080")
	defer hcServer.Shutdown()

	// run
	log.Println("startup complete")
	<-stop

	// shutdown
	log.Println("shutting down")
	cancel()
}
