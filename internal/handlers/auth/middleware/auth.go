package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/mortenoj/go-graphql-template/pkg/utils"
)

func authError(c *gin.Context, err error) {
	errKey := "message"
	errMsgHeader := "[Auth] error: "
	e := gin.H{errKey: errMsgHeader + err.Error()}
	c.AbortWithStatusJSON(http.StatusUnauthorized, e)
}

// Middleware wraps the request with auth middleware
func Middleware(path string, cfg *config.Config, orm *orm.ORM) gin.HandlerFunc {
	logrus.Info("[Auth.Middleware] Applied to path: ", path)

	return gin.HandlerFunc(func(c *gin.Context) {
		a, err := ParseAPIKey(c, cfg)
		if err != nil {
			if err != ErrEmptyAPIKeyHeader {
				authError(c, err)
				return
			}

			t, err := ParseToken(c, cfg)
			if err != nil {
				authError(c, err)
				return
			}

			claims, ok := t.Claims.(jwt.MapClaims)
			if !ok {
				authError(c, err)
				return
			}

			if claims["exp"] == nil {
				authError(c, ErrMissingExpField)
				return
			}

			issuer := claims["iss"].(string)
			userid := claims["jti"].(string)
			email := claims["email"].(string)

			user, err := orm.FindUserByJWT(email, issuer, userid)
			if err != nil {
				authError(c, ErrForbidden)
				return
			}

			c.Request = addToContext(c, utils.ProjectContextKeys().UserCtxKey, user)
			c.Next()
			return
		}

		_, err = orm.FindUserByAPIKey(a)
		if err != nil {
			authError(c, ErrForbidden)
			return
		}

		c.Next()
	})
}
