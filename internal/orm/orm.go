// Package orm handles ORM operations
package orm

import (
	"github.com/mortenoj/reko-ring-backend/internal/config"
	"github.com/mortenoj/reko-ring-backend/internal/orm/migrations"
	"github.com/sirupsen/logrus"

	// Dialect also imports postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func Factory(cfg *config.Config) (*ORM, error) {
	db, err := gorm.Open(cfg.Gorm.Dialect, cfg.Gorm.ConnectionDSN)
	if err != nil {
		logrus.Error("[ORM] err: ", err)
		return nil, err
	}

	orm := &ORM{
		DB: db,
	}

	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(cfg.Gorm.LogMode)

	// Automigrate tables
	if cfg.Gorm.AutoMigrate {
		err = migrations.ServiceAutoMigrations(orm.DB)
	}

	logrus.Info("[ORM] Database connection initialized.")

	return orm, err
}
