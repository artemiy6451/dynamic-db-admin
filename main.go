// Package main is the entry point for Dynamic DB Admin.
package main

import (
	"embed"
	"log"

	"artemiy6451/dynamic-db-admin/internal/app"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/logging"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	application.RegisterEvent[string]("time")
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load application configuration: %v", err)
	}

	logger := logging.New(logging.Config{
		Development: cfg.Logging.Development,
		Level:       cfg.Logging.Level,
		Output:      cfg.Logging.Output,
	})

	application, err := app.New(logger, cfg, assets)
	if err != nil {
		log.Fatalf("initialize application: %v", err)
	}
	if err := application.Run(); err != nil {
		log.Fatalf("run application: %v", err)
	}
}
