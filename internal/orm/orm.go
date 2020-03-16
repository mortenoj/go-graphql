// Package orm handles ORM operations
package orm

import (
	"errors"

	"github.com/markbates/goth"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/gql/resolvers/transformations"
	"github.com/mortenoj/go-graphql-template/internal/orm/migrations"
	"github.com/mortenoj/go-graphql-template/internal/orm/models"

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
	db, err := gorm.Open(cfg.Database.Dialect, cfg.Database.ConnectionDSN)
	if err != nil {
		logrus.Error("[ORM] err: ", err)
		return nil, err
	}

	orm := &ORM{
		DB: db,
	}

	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(cfg.Database.LogMode)

	// Automigrate tables
	if cfg.Database.AutoMigrate {
		err = migrations.ServiceAutoMigrations(orm.DB)
	}

	logrus.Info("[ORM] Database connection initialized.")

	return orm, err
}

//FindUserByAPIKey finds the user that is related to the API key
func (o *ORM) FindUserByAPIKey(apiKey string) (*models.User, error) {
	db := o.DB.New()
	uak := &models.UserAPIKey{}

	if apiKey == "" {
		return nil, errors.New("API key is empty")
	}

	if err := db.Preload("User").Where("api_key = ?", apiKey).Find(uak).Error; err != nil {
		return nil, err
	}

	return &uak.User, nil
}

// FindUserByJWT finds the user that is related to the APIKey token
func (o *ORM) FindUserByJWT(email string, provider string, userID string) (*models.User, error) {
	db := o.DB.New()
	up := &models.UserProfile{}

	if provider == "" || userID == "" {
		return nil, errors.New("provider or userId empty")
	}

	err := db.Preload("User").Where(
		"email  = ? AND provider = ? AND external_user_id = ?",
		email,
		provider,
		userID,
	).First(up).Error
	if err != nil {
		return nil, err
	}

	return &up.User, nil
}

// UpsertUserProfile saves the user if doesn't exists and adds the OAuth profile
func (o *ORM) UpsertUserProfile(input *goth.User) (*models.User, error) {
	db := o.DB.New()
	up := &models.UserProfile{}

	u, err := transformations.GothUserToDBUser(input, false)
	if err != nil {
		return nil, err
	}

	if tx := db.Where("email = ?", input.Email).First(u); !tx.RecordNotFound() && tx.Error != nil {
		return nil, tx.Error
	}

	if tx := db.Model(u).Save(u); tx.Error != nil {
		return nil, err
	}

	if tx := db.Where("email = ? AND provider = ? AND external_user_id = ?",
		input.Email, input.Provider, input.UserID).First(up); !tx.RecordNotFound() && tx.Error != nil {
		return nil, err
	}

	up, err = transformations.GothUserToDBUserProfile(input, false)
	if err != nil {
		return nil, err
	}

	up.User = *u
	if tx := db.Model(up).Save(up); tx.Error != nil {
		return nil, tx.Error
	}

	logrus.Info(u.ID)

	return u, nil
}
