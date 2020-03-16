// Package jobs has seed jobs
package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/mortenoj/go-graphql-template/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

func users() []*models.User {
	var (
		uname       = "Test User"
		fname       = "Test"
		lname       = "User"
		nname       = "Foo Bar"
		description = "This is the first user ever!"
		location    = "His house, maybe?"
	)

	return []*models.User{
		{
			Email:       "test@test.com",
			Name:        &uname,
			FirstName:   &fname,
			LastName:    &lname,
			NickName:    &nname,
			Description: &description,
			Location:    &location,
		},
		{
			Email:       "test2@test.com",
			Name:        &uname,
			FirstName:   &fname,
			LastName:    &lname,
			NickName:    &nname,
			Description: &description,
			Location:    &location,
		},
	}
}

// SeedUsers defines job for seeding users table
func SeedUsers() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "SEED_USERS",
		Migrate: func(db *gorm.DB) error {
			for _, u := range users() {
				if err := db.Create(u).Error; err != nil {
					return err
				}
				if err := db.Create(&models.UserAPIKey{UserID: u.ID}).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			for _, u := range users() {
				if err := db.Delete(u).Error; err != nil {
					return err
				}
			}
			return nil
		},
	}
}
