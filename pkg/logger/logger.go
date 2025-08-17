package logger

import (
	"io"
	"log/slog"
	"os"
)

var (
	// LevelFlagOptions represents allowed logging levels.
	LevelFlagOptions = []string{"debug", "info", "warn", "error"}
	// FormatFlagOptions represents allowed formats.
	FormatFlagOptions = []string{"logfmt", "json"}

	defaultWriter = os.Stdout
)

type Level struct {
	dLevel *slog.LevelVar
}

type Config struct {
	Level  *Level
	Format string
	Style  string
	Writer io.Writer
}

func NewLevel() *Level {
	return &Level{
		dLevel: &slog.LevelVar{},
	}
}

func New(config *Config) *slog.Logger {
	if config.Level == nil {
		config.Level = NewLevel()
	}

	if config.Writer == nil {
		config.Writer = defaultWriter
	}

	logHandlerOpts := &slog.HandlerOptions{
		Level:     config.Level.dLevel,
		AddSource: true,
	}

	if config.Format != "" && config.Format == "json" {
		return slog.New(slog.NewJSONHandler(config.Writer, logHandlerOpts))
	}
	return slog.New(slog.NewTextHandler(config.Writer, logHandlerOpts))
}
