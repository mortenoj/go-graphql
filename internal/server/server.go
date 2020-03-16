// Package server holds all API operations
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/handlers"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/mortenoj/go-graphql-template/pkg/utils/errutils"
	"github.com/sirupsen/logrus"
)

// Run web server
func Run(cfg *config.Config, orm *orm.ORM) {
	logrus.Info("GORM_CONNECTION_DSN: ", cfg.Gorm.ConnectionDSN)

	r := gin.Default()

	gin.SetMode(cfg.GinMode)

	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL and Playground handler
	if cfg.GraphQL.PlaygroundEnabled {
		r.GET(cfg.GraphQL.PlaygroundPath, handlers.PlaygroundHandler(cfg.GraphQL.ServerPath))
	}

	r.POST(cfg.GraphQL.ServerPath, handlers.GraphqlHandler(orm))

	errutils.Must(r.Run(cfg.GraphQL.Host + ":" + cfg.GraphQL.Port))
}
