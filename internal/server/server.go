// Package server holds all API operations
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/google"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/sirupsen/logrus"
)

// Run web server
func Run(cfg *config.Config, orm *orm.ORM) error {
	var err error

	r := gin.Default()

	// Initialize the Auth providers
	err = InitalizeAuthProviders(cfg)
	if err != nil {
		return err
	}

	// Routes and Handlers
	err = RegisterRoutes(cfg, r, orm)
	if err != nil {
		return err
	}

	// Inform the user where the server is listening
	logrus.Info("Running @ " + cfg.Server.SchemaVersionedEndpoint(""))

	// Run the server
	// Print out and exit(1) to the OS if the server cannot run
	return r.Run(cfg.Server.ListenEndpoint())
}

// InitalizeAuthProviders does just that, with Goth providers
func InitalizeAuthProviders(cfg *config.Config) error {
	providers := []goth.Provider{}
	// Initialize Goth providers
	if cfg.AuthProviders.Google.Key != "" {
		providers = append(providers, google.New(cfg.AuthProviders.Google.Key, cfg.AuthProviders.Google.Secret,
			cfg.SchemaVersionedEndpoint("/auth/google/callback"),
			cfg.AuthProviders.Google.Scopes...))
	}

	if cfg.AuthProviders.Auth0.Key != "" {
		providers = append(providers, auth0.New(cfg.AuthProviders.Auth0.Key, cfg.AuthProviders.Auth0.Secret,
			cfg.SchemaVersionedEndpoint("/auth/auth0/callback"),
			cfg.AuthProviders.Auth0.Domain, cfg.AuthProviders.Auth0.Scopes...))
	}

	goth.UseProviders(providers...)

	return nil
}
