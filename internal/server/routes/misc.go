package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/handlers"
	"github.com/mortenoj/go-graphql-template/internal/handlers/auth/middleware"
	"github.com/mortenoj/go-graphql-template/internal/orm"
)

// Misc routes
func Misc(cfg *config.Config, r *gin.Engine, orm *orm.ORM) error {
	// Simple keep-alive/ping handler
	r.GET(cfg.VersionedEndpoint("/ping"), handlers.Ping())
	r.GET(cfg.VersionedEndpoint("/secure-ping"),
		middleware.Middleware(cfg.Server.VersionedEndpoint("/secure-ping"), cfg, orm), handlers.Ping())

	return nil
}
