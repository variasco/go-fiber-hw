package logger

import (
	"os"
	"strings"
	"variasco/go-fiber-hw/config"

	"github.com/rs/zerolog"
)

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	level := parseLogLevel(config.Level)
	zerolog.SetGlobalLevel(level)
	logger := parseLogFormat(config.Format)

	return &logger
}

func parseLogFormat(logFormat string) zerolog.Logger {
	switch strings.ToLower(logFormat) {
	case "console":
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		return zerolog.New(consoleWriter).With().Timestamp().Logger()
	case "json":
		return zerolog.New(os.Stdout).With().Timestamp().Logger()
	default:
		return zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case zerolog.LevelTraceValue:
		return zerolog.TraceLevel
	case zerolog.LevelDebugValue:
		return zerolog.DebugLevel
	case zerolog.LevelInfoValue:
		return zerolog.InfoLevel
	case zerolog.LevelWarnValue:
		return zerolog.WarnLevel
	case zerolog.LevelErrorValue:
		return zerolog.ErrorLevel
	case zerolog.LevelFatalValue:
		return zerolog.FatalLevel
	case zerolog.LevelPanicValue:
		return zerolog.PanicLevel
	default:
		return zerolog.Disabled
	}
}
