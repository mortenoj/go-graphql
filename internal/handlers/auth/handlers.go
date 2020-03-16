package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/sirupsen/logrus"
)

// Claims JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Begin login with the auth provider
func Begin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// You have to add value context with provider name to get provider name in GetProviderName method
		c.Request = addProviderToContext(c, c.Param("provider"))
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err != nil {
			gothic.BeginAuthHandler(c.Writer, c.Request)
		} else {
			logrus.Infof("user: %#v", gothUser)
		}
	}
}

// Callback callback to complete auth provider flow
func Callback(cfg *config.Config, orm *orm.ORM) gin.HandlerFunc {
	return func(c *gin.Context) {
		// You have to add value context with provider name to get provider name in GetProviderName method
		c.Request = addProviderToContext(c, c.Param("provider"))

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		u, err := orm.FindUserByJWT(user.Email, user.Provider, user.UserID)
		// logrus.Infof("gothUser: %#v", user)
		if err != nil {
			if u, err = orm.UpsertUserProfile(&user); err != nil {
				logrus.Errorf("[Auth.CallBack.UserLoggedIn.UpsertUserProfile.Error]: %q", err)
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
		}

		logrus.Info("[Auth.CallBack.UserLoggedIn]: ", u)

		jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.JWT.Algorithm), Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				Id:        user.UserID,
				Issuer:    user.Provider,
				IssuedAt:  time.Now().UTC().Unix(),
				NotBefore: time.Now().UTC().Unix(),
				ExpiresAt: user.ExpiresAt.UTC().Unix(),
			},
		})

		token, err := jwtToken.SignedString([]byte(cfg.JWT.Secret))
		if err != nil {
			logrus.Error("[Auth.Callback.JWT] error: ", err)
			_ = c.AbortWithError(http.StatusInternalServerError, err)

			return
		}

		logrus.Info("token: ", token)

		json := gin.H{
			"type":          "Bearer",
			"token":         token,
			"refresh_token": user.RefreshToken,
		}

		c.JSON(http.StatusOK, json)
	}
}

// Logout logs out of the auth provider
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = addProviderToContext(c, c.Param("provider"))

		err := gothic.Logout(c.Writer, c.Request)
		if err != nil {
			logrus.Error("[gothic] error logging out")
		}

		c.Writer.Header().Set("Location", "/")
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	}
}
