// Package config for import config file to project
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigService struct {
	config Config
}

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Window      WindowConfig `yaml:"window"`
}

type WindowConfig struct {
	Title     string `yaml:"title"`
	MinWidth  int    `yaml:"min_width"`
	MinHeight int    `yaml:"min_height"`
	Width     int    `yaml:"width"`
	Height    int    `yaml:"height"`
}

func New() Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Can not read config file with err: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Can not parce config file with err: %s", err)
	}

	log.Printf("Config file parced: %v+\n", config)

	return config
}

func NewConfigService(config Config) *ConfigService {
	return &ConfigService{
		config: config,
	}
}

func (s *ConfigService) GetConfig() Config {
	return s.config
}
