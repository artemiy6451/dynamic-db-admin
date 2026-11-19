// Package main for sturtup application
package main

import (
	"log"

	"artemiy6451/dynamic-db-admin/internal/app"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/logging"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

func main() {
	cfg := config.New()
	logger := logging.New(
		logging.Config{
			Development: cfg.Logging.Development,
			Level:       cfg.Logging.Level,
			Output:      cfg.Logging.Output,
		})

	app, err := app.New(logger, cfg)
	if err != nil {
		log.Fatalf("application failed: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}
