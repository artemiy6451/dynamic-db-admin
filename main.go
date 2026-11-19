// Package main for sturtup application
package main

import (
	"log"
	"time"

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

	app := app.New(logger, cfg)

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()
	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
