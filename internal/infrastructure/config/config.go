// Package config loads application settings from the user's configuration directory.
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"artemiy6451/dynamic-db-admin/internal/infrastructure/storage"
	"gopkg.in/yaml.v3"
)

type ConfigService struct {
	config Config
}

type Config struct {
	App     AppConfig     `yaml:"app"`
	Logging LoggingConfig `yaml:"logging"`
}

type AppConfig struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Version     string       `yaml:"version"`
	Window      WindowConfig `yaml:"window"`
}

type WindowConfig struct {
	Title     string `yaml:"title"`
	MinWidth  int    `yaml:"min_width"`
	MinHeight int    `yaml:"min_height"`
	Width     int    `yaml:"width"`
	Height    int    `yaml:"height"`
}

type LoggingConfig struct {
	Development bool   `yaml:"development"`
	Level       string `yaml:"level"`
	Output      string `yaml:"output"`
}

func Default() Config {
	return Config{
		App: AppConfig{
			Name:        "Dynamic DB Admin",
			Description: "Desktop application for visual SQLite database management",
			Version:     "0.0.1",
			Window: WindowConfig{
				Title: "Dynamic DB Admin", MinWidth: 1000, MinHeight: 618,
				Width: 1000, Height: 618,
			},
		},
		Logging: LoggingConfig{
			Development: false,
			Level:       "INFO",
			Output:      "std",
		},
	}
}

// New loads an optional YAML configuration. By default it reads from the
// per-user configuration directory, never from the executable's directory.
// DYNAMIC_DB_ADMIN_CONFIG can be set to an explicit path for development.
func New() (*Config, error) {
	cfg := Default()

	configPath := os.Getenv("DYNAMIC_DB_ADMIN_CONFIG")
	if configPath == "" {
		paths, err := storage.ResolvePaths()
		if err != nil {
			return nil, fmt.Errorf("resolve application storage paths: %w", err)
		}
		configPath = filepath.Join(paths.ConfigDir, "config.yaml")
	}

	data, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return &cfg, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", configPath, err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", configPath, err)
	}
	if cfg.App.Name == "" || cfg.App.Window.Width <= 0 || cfg.App.Window.Height <= 0 {
		return nil, fmt.Errorf("configuration %q contains invalid application metadata or window dimensions", configPath)
	}

	return &cfg, nil
}

func NewConfigService(config Config) *ConfigService {
	return &ConfigService{config: config}
}

func (s *ConfigService) GetConfig() Config {
	return s.config
}
