package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/handlers"
	auth "github.com/mortenoj/go-graphql-template/internal/handlers/auth/middleware"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/sirupsen/logrus"
)

// GraphQL routes
func GraphQL(cfg *config.Config, r *gin.Engine, orm *orm.ORM) error {
	// GraphQL paths
	gqlPath := cfg.Server.VersionedEndpoint(cfg.GraphQL.ServerPath)
	pgqlPath := cfg.GraphQL.PlaygroundPath
	g := r.Group(gqlPath)

	// GraphQL handler
	g.POST("", auth.Middleware(g.BasePath(), cfg, orm), handlers.GraphqlHandler(orm))
	logrus.Info("GraphQL @ ", gqlPath)

	// Playground handler
	if cfg.GraphQL.PlaygroundEnabled {
		logrus.Info("GraphQL Playground @ ", g.BasePath()+pgqlPath)
		g.GET(pgqlPath, handlers.PlaygroundHandler(g.BasePath()))
	}

	return nil
}
