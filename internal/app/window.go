package app

import (
	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func createMainWindow(app *application.App, cfg *config.Config) {
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            cfg.App.Window.Title,
		MinWidth:         cfg.App.Window.MinWidth,
		MinHeight:        cfg.App.Window.MinHeight,
		Width:            cfg.App.Window.Width,
		Height:           cfg.App.Window.Height,
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})
}
