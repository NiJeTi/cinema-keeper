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
)

type config struct {
	Discord discord.Config `conf:"discord"`
	DB      db.Config      `conf:"db"`
}

func main() {
	log.Println("starting")

	cfg := cfgPkg.ReadConfig[config]()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")
	cmdLogger := logger.WithGroup("command")

	// db
	dbConn := db.Connect(cfg.DB)
	defer dbConn.Close()
	txWrapper := dbUtils.NewTxWrapper(dbLogger, dbConn)

	// discord
	discordConn := discord.Connect(cfg.Discord.Token)
	defer discordConn.Close()

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

	// run
	log.Println("started")
	<-stop
	log.Println("shutting down")
}
