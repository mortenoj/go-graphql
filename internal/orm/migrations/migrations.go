// Package migrations handles migrations
package migrations

import (
	"fmt"

	"github.com/mortenoj/go-graphql-template/internal/orm/migrations/jobs"
	"github.com/mortenoj/go-graphql-template/internal/orm/models"
	"github.com/mortenoj/go-graphql-template/pkg/utils/consts"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

const logPrefix string = "[migrations] "

// ServiceAutoMigrations migrates all the tables and modifications to the connected source
func ServiceAutoMigrations(db *gorm.DB) error {
	var err error

	// Initialize the migration empty so InitSchema runs always first on creation
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		logrus.Info("[Migration.InitSchema] Initializing database schema")

		if db.Dialect().GetName() == "postgres" {
			db.Exec("CREATE EXTENSION IF NOT EXISTS\"uuid-ossp\";")
		}

		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}

		return nil
	})

	err = m.Migrate()
	if err != nil {
		logrus.Error(logPrefix, "error running initial migrations")
		return err
	}

	if err := updateMigration(db); err != nil {
		return err
	}

	// Keep a list of migrations here
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers(),
		jobs.SeedRBAC(),
	})

	err = m.Migrate()
	if err != nil {
		logrus.Error(logPrefix, "error running initial migrations")
		return err
	}

	return nil
}

func updateMigration(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.UserProfile{},
		&models.UserAPIKey{},
	).Error
	if err != nil {
		return err
	}

	return addIndexes(db)
}

func addIndexes(db *gorm.DB) (err error) {
	// Entity names
	//db.NewScope(&models.User{}).GetModelStruct().TableName(db)
	usersTableName := consts.Tablenames().Users
	rolesTableName := consts.Tablenames().Roles
	permissionsTableName := consts.Tablenames().Permissions

	// FKs
	if err := db.Model(&models.UserProfile{}).
		AddForeignKey("user_id", usersTableName+"(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}

	if err := db.Model(&models.UserAPIKey{}).
		AddForeignKey("user_id", usersTableName+"(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}

	if err := db.Model(&models.UserRole{}).
		AddForeignKey("user_id", usersTableName+"(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}

	if err := db.Model(&models.UserRole{}).
		AddForeignKey("role_id", rolesTableName+"(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}

	if err := db.Model(&models.UserPermission{}).
		AddForeignKey("user_id", usersTableName+"(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}

	if err := db.Model(&models.UserPermission{}).
		AddForeignKey("permission_id", permissionsTableName+"(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	// Indexes
	// None needed so far
	return nil
}
