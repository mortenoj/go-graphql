package config

// Database holds db configs
type Database struct {
	Dialect       string `env:"GORM_DIALECT"`
	ConnectionDSN string `env:"GORM_CONNECTION_DSN"`
	SeedDB        bool   `env:"GORM_SEED_DB"`
	LogMode       bool   `env:"GORM_LOGMODE"`
	AutoMigrate   bool   `env:"GORM_AUTOMIGRATE"`
}
