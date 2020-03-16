// Package routes handles server routes
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/handlers/auth"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/mortenoj/go-graphql-template/pkg/utils"
)

// Auth routes
func Auth(cfg *config.Config, r *gin.Engine, orm *orm.ORM) error {
	provider := string(utils.ProjectContextKeys().ProviderCtxKey)

	// OAuth handlers
	g := r.Group(cfg.Server.VersionedEndpoint("/auth"))
	g.GET("/:"+provider, auth.Begin())
	g.GET("/:"+provider+"/callback", auth.Callback(cfg, orm))
	// g.GET(:"+provider+"/refresh", auth.Refresh(cfg, orm))

	return nil
}
