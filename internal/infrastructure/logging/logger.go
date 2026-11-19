// Package logging setup logging using tint for development and Json for production
package logging

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type Config struct {
	Development bool
	Level       string
	Output      string
}

func New(cfg Config) *slog.Logger {
	level := parceLevel(cfg.Level)
	output := parceOutput(cfg.Output)

	if output == nil {
		output = os.Stdout
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.Development,
	}

	var handler slog.Handler

	if cfg.Development {
		handler = tint.NewTextHandler(output, &tint.Options{
			Level:      level,
			AddSource:  true,
			TimeFormat: "15:04:05",
		})
	} else {
		handler = slog.NewJSONHandler(output, opts)
	}

	return slog.New(handler)
}
