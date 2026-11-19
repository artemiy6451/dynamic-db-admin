// Package app creates and configures the Wails application.
package app

import (
	"embed"
	"log/slog"

	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	wails *application.App
	deps  *Dependencies
}

func New(logger *slog.Logger, cfg *config.Config, assets embed.FS) (*App, error) {
	wailsApp, deps, err := Bootstrap(cfg, logger, assets)
	if err != nil {
		return nil, err
	}
	return &App{
		wails: wailsApp,
		deps:  deps,
	}, nil
}

func (a *App) Run() error {
	a.deps.Logger.Info("starting application")
	return a.wails.Run()
}
