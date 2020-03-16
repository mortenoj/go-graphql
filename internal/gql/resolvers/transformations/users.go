// Package transformations converts db models to gql models
package transformations

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/markbates/goth"
	gql "github.com/mortenoj/go-graphql-template/internal/gql/models"
	dbm "github.com/mortenoj/go-graphql-template/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		AvatarURL:   i.AvatarURL,
		ID:          i.ID.String(),
		Email:       i.Email,
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}

	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	o = &dbm.User{
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
	}

	if i.Email == nil && !update {
		return nil, errors.New("field [email] is required")
	}

	if i.Password == nil && !update {
		return nil, errors.New("field [password] is required")
	}

	if i.Email != nil {
		o.Email = *i.Email
	}

	if i.Password != nil {
		o.Password = *i.Password
	}

	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}

		o.ID = updID
	}

	return o, err
}

// GothUserToDBUser transforms [user] goth to db model
func GothUserToDBUser(i *goth.User, update bool, ids ...string) (o *dbm.User, err error) {
	if i.Email == "" && !update {
		return nil, errors.New("field [Email] is required")
	}

	o = &dbm.User{
		Email:       i.Email,
		Name:        &i.Name,
		FirstName:   &i.FirstName,
		LastName:    &i.LastName,
		NickName:    &i.NickName,
		Location:    &i.Location,
		AvatarURL:   &i.AvatarURL,
		Description: &i.Description,
	}

	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}

		o.ID = updID
	}

	return o, err
}

// GothUserToDBUserProfile transforms [user] goth to db model
func GothUserToDBUserProfile(i *goth.User, update bool, ids ...int) (o *dbm.UserProfile, err error) {
	if i.UserID == "" && !update {
		return nil, errors.New("field [UserID] is required")
	}

	if i.Email == "" && !update {
		return nil, errors.New("field [Email] is required")
	}

	o = &dbm.UserProfile{
		ExternalUserID: i.UserID,
		Provider:       i.Provider,
		Email:          i.Email,
		Name:           i.Name,
		FirstName:      i.FirstName,
		LastName:       i.LastName,
		NickName:       i.NickName,
		Location:       i.Location,
		AvatarURL:      i.AvatarURL,
		Description:    &i.Description,
	}

	if len(ids) > 0 {
		updID := ids[0]
		o.ID = updID
	}

	return o, err
}
