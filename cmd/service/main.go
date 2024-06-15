package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"github.com/nijeti/cinema-keeper/internal/commands"
	"github.com/nijeti/cinema-keeper/internal/db"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/quote"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
	"github.com/nijeti/cinema-keeper/internal/pkg/config"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbUtils"
	"github.com/nijeti/cinema-keeper/internal/types"
)

func main() {
	config.Config()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbLogger := logger.WithGroup("db")
	cmdLogger := logger.WithGroup("command")

	// db
	dbConn := config.DB(viper.GetString("db.connection_string"))
	defer dbConn.Close()
	txWrapper := dbUtils.NewTxWrapper(dbLogger, dbConn)

	// discord
	discord := config.Connect(viper.GetString("discord.token"))
	defer discord.Close()

	// repos
	quotesRepo := db.NewQuotesRepo(dbLogger, dbConn, txWrapper)

	// handlers
	cmds := map[string]*commands.Command{
		commands.QuoteName: {
			Description: commands.Quote,
			Handler:     quote.New(ctx, cmdLogger, quotesRepo),
		},
		commands.CastName: {
			Description: commands.Cast,
			Handler:     cast.New(ctx, cmdLogger),
		},
		commands.LockName: {
			Description: commands.Lock,
			Handler:     lock.New(ctx, cmdLogger),
		},
		commands.UnlockName: {
			Description: commands.Unlock,
			Handler:     unlock.New(ctx, cmdLogger),
		},
	}

	// handlers
	guildID := types.ID(viper.GetString("discord.guild"))
	config.RegisterCommands(discord, cmds, guildID)
	defer config.UnregisterCommands(discord, cmds, guildID)

	// run
	<-stop
}
