// Package migrations handles migrations
package migrations

import (
	"fmt"

	"github.com/mortenoj/go-graphql-template/internal/gql/models"
	"github.com/mortenoj/go-graphql-template/internal/orm/migrations/jobs"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

const logPrefix string = "[migrations] "

// ServiceAutoMigrations migrates all the tables and modifications to the connected source
func ServiceAutoMigrations(db *gorm.DB) error {
	var err error

	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		logrus.Info("[Migration.InitSchema] Initializing database schema")

		if db.Dialect().GetName() == "postgres" {
			// Let's create the UUID extension, the user has to have superuser permission for now
			db.Exec("create extension \"uuid-ossp\";")
		}
		if err := updateMigrations(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}

		return nil
	})

	err = m.Migrate()
	if err != nil {
		logrus.Error(logPrefix, "error running init migrations")
		return err
	}

	if err := updateMigrations(db); err != nil {
		logrus.Error(logPrefix, "error updating migrations")
		return err
	}

	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers(),
	})

	err = m.Migrate()
	if err != nil {
		logrus.Error(logPrefix, "error running updated migrations")
		return err
	}

	return nil
}

func updateMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}
