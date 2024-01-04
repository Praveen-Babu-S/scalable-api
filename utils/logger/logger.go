package logger

import (
	"log/slog"
	"os"
)

func CreateLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}
