// Package auth route
package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/pkg/utils"
)

func addProviderToContext(c *gin.Context, value interface{}) *http.Request {
	//nolint (ProviderCtxKey must be wrapped in string())
	return c.Request.WithContext(
		context.WithValue(c.Request.Context(),
			string(utils.ProjectContextKeys().ProviderCtxKey),
			value,
		))
}
