package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/mortenoj/go-graphql-template/internal/server/routes"
)

// RegisterRoutes register the routes for the server
func RegisterRoutes(cfg *config.Config, r *gin.Engine, orm *orm.ORM) (err error) {
	// Auth routes
	if err = routes.Auth(cfg, r, orm); err != nil {
		return err
	}

	// GraphQL server routes
	if err = routes.GraphQL(cfg, r, orm); err != nil {
		return err
	}

	// Miscellaneous routes
	if err = routes.Misc(cfg, r, orm); err != nil {
		return err
	}

	return err
}
