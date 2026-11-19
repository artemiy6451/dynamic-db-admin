package logging

import (
	"io"
	"log/slog"
	"os"
)

func parceLevel(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func parceOutput(output string) io.Writer {
	switch output {
	case "std":
		return os.Stdout
	// case "file": TODO: add file output
	default:
		return os.Stdout
	}
}
