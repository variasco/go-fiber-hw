package logger

import (
	"log/slog"
	"os"
	"strings"
	"variasco/go-fiber-hw/config"
)

func CreateLogger(config *config.LogConfig) *slog.Logger {
	level := parseLogLevel(config.Level)
	handler := parseLogFormat(config.Format, &slog.HandlerOptions{Level: level})

	return slog.New(handler)
}

func parseLogFormat(logFormat string, opts *slog.HandlerOptions) slog.Handler {
	switch strings.ToLower(logFormat) {
	case "console":
		return slog.NewTextHandler(os.Stdout, opts)
	case "json":
		return slog.NewJSONHandler(os.Stderr, opts)
	default:
		return slog.NewJSONHandler(os.Stdout, opts)
	}
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
