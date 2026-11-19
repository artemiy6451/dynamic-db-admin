// Package app for create and setup wails application
package app

import (
	"embed"
	"log/slog"

	wailsadapter "artemiy6451/dynamic-db-admin/internal/adapters/wails"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func New(logger *slog.Logger, cfg *config.Config) *application.App {
	var assets embed.FS
	appService := wailsadapter.NewAppService(logger, cfg)

	app := application.New(application.Options{
		Name:        cfg.App.Name,
		Description: cfg.App.Description,
		Services: []application.Service{
			application.NewService(appService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	createMainWindow(app, cfg)
	return app
}
