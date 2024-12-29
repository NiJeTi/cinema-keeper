package main

import (
	"fmt"
	"log/slog"
)

type gooseLogger struct {
	logger *slog.Logger
}

func (l gooseLogger) Fatalf(format string, v ...any) {
	l.logger.Error(fmt.Sprintf(format, v...)) //nolint:sloglint // adapter
}

func (l gooseLogger) Printf(format string, v ...any) {
	l.logger.Info(fmt.Sprintf(format, v...)) //nolint:sloglint // adapter
}
