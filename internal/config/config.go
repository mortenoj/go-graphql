// Package config reads and handles environment variables and configs
package config

import (
	"github.com/caarlos0/env/v6"
)

// Gorm holds GORM configs
type Gorm struct {
	Dialect       string `env:"GORM_DIALECT"`
	ConnectionDSN string `env:"GORM_CONNECTION_DSN"`
	SeedDB        bool   `env:"GORM_SEED_DB"`
	LogMode       bool   `env:"GORM_LOGMODE"`
	AutoMigrate   bool   `env:"GORM_AUTOMIGRATE"`
}

// GraphQL holds server configurations
type GraphQL struct {
	Host              string `env:"GQL_SERVER_HOST" envDefault:"localhost"`
	Port              string `env:"GQL_SERVER_PORT" envDefault:"8080"`
	ServerPath        string `env:"GQL_SERVER_GRAPHQL_PATH" envDefault:"/graphql"`
	PlaygroundEnabled bool   `env:"GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED" envDefault:"true"`
	PlaygroundPath    string `env:"GQL_SERVER_GRAPHQL_PLAYGROUND_PATH" envDefault:"/"`
}

// Config holds the application configs
type Config struct {
	GraphQL
	Gorm
	GinMode    string `env:"GIN_MODE" envDefault:"debug"`
	Production bool   `env:"PRODUCTION"`
}

// Init initializes the config values
func Init() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
