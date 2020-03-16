// Package config reads and handles environment variables and configs
package config

import (
	"github.com/caarlos0/env/v6"
)

// Config holds the application configs
type Config struct {
	Server
	GraphQL
	Database
	JWT
	AuthProviders
	Production bool `env:"PRODUCTION"`
}

// Init initializes the config values
func Init() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
