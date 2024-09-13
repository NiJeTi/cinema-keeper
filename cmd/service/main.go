package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/nijeti/healthcheck"
	"github.com/nijeti/healthcheck/servers/fasthttp"

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
	dbConn := db.Connect(cfg.DB.ConnectionString)
	defer dbConn.Close()
	dbProbe := db.NewProbe(dbConn)
	txWrapper := dbUtils.NewTxWrapper(dbLogger, dbConn)

	quotesRepo := db.NewQuotesRepo(dbLogger, dbConn, txWrapper)

	// discord
	discordSession := discord.Connect(cfg.Discord.Token)
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
	discord.RegisterCommands(discordSession, commands, cfg.Discord.Guild)
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
}
