package logger

import (
	"log/slog"
	"os"
	"variasco/go-fiber-hw/config"
)

func CreateLogger(config *config.LogConfig) *slog.Logger {
	if config.Format == "console" {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		return slog.New(slog.NewJSONHandler(os.Stderr, nil))
	}
}
