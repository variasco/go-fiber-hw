package logger

import (
	"log/slog"
	"os"
	"strings"
	"variasco/go-fiber-hw/config"
)

func CreateLogger(config *config.LogConfig) *slog.Logger {
	level := parseLogLevel(config.Level)
	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if config.Format == "console" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	return slog.New(handler)
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
