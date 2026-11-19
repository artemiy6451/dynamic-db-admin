package wails

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"artemiy6451/dynamic-db-admin/internal/infrastructure/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AppService struct {
	logger *slog.Logger
	cfg    *config.Config

	mu  sync.Mutex
	ctx context.Context
}

func NewAppService(logger *slog.Logger, cfg *config.Config) *AppService {
	return &AppService{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *AppService) ServiceStartup(ctx context.Context, options application.Options) error {
	s.mu.Lock()
	s.ctx = ctx
	s.mu.Unlock()

	s.logger.Info("Application service started", "component", "app_service")

	return nil
}

func (s *AppService) ServiceShutdown() error {
	s.logger.Info("Application service stopped", "component", "app_service")

	s.mu.Lock()
	s.ctx = nil
	s.mu.Unlock()

	return nil
}

func (s *AppService) GetAppInfo() AppInfo {
	return AppInfo{
		Name:    s.cfg.App.Name,
		Version: s.cfg.App.Version,
	}
}

func (s *AppService) Ping() (string, error) {
	s.mu.Lock()
	ctx := s.ctx

	s.mu.Unlock()

	if ctx == nil {
		return "", fmt.Errorf("application service is not initialized")
	}

	if err := ctx.Err(); err != nil {
		return "", fmt.Errorf("application is shutting down: %w", err)
	}

	return "Backend is running", nil
}

func (s *AppService) Greet(name string) string {
	name = strings.TrimSpace(name)

	if name == "" {
		name = "Developer"
	}

	s.logger.Debug("greetings request")

	return fmt.Sprintf("Hello, %s! Welcome to Dynamic DB Admin.", name)
}
