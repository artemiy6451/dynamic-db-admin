package app

import (
	"embed"
	"log/slog"

	wailsadapter "artemiy6451/dynamic-db-admin/internal/adapters/wails"
	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Dependencies struct {
	Logger     *slog.Logger
	AppService *wailsadapter.AppService
}

// InitializeDependencies constructs the application's dependencies.
func InitializeDependencies(cfg *config.Config, logger *slog.Logger) (*Dependencies, error) {
	appService := wailsadapter.NewAppService(logger, cfg)
	return &Dependencies{
		Logger:     logger,
		AppService: appService,
	}, nil
}

// Bootstrap creates and configures the Wails application.
func Bootstrap(cfg *config.Config, logger *slog.Logger, assets embed.FS) (*application.App, *Dependencies, error) {
	deps, err := InitializeDependencies(cfg, logger)
	if err != nil {
		return nil, nil, err
	}

	wailsApp := application.New(application.Options{
		Name:        cfg.App.Name,
		Description: cfg.App.Description,
		Services: []application.Service{
			application.NewService(deps.AppService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	createMainWindow(wailsApp, cfg)
	deps.Logger.Info("application initialized")
	return wailsApp, deps, nil
}
