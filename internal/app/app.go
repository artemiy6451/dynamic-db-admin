// Package app for create and setup wails application
package app

import (
	"log/slog"

	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	wails *application.App
	dpes  *Dependencies
}

func New(logger *slog.Logger, cfg *config.Config) (*App, error) {
	wailsApp, deps, err := Bootstrap(cfg, logger)
	if err != nil {
		return nil, err
	}
	return &App{
		wails: wailsApp,
		dpes:  deps,
	}, nil
}

func (a *App) Run() error {
	a.dpes.Logger.Info("starting application")

	return a.wails.Run()
}
