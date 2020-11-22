package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

type AuthMw struct {
	ac auth.AuthServiceClient
}

func NewAuthMw(ac auth.AuthServiceClient) *AuthMw {
	return &AuthMw{
		ac: ac,
	}
}

func (mw *AuthMw) AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(config.Conf.Token.CookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "bad cookie"})
			return
		}

		userId, err := mw.ac.CheckCookie(c, &auth.Check{Cookie: cookie})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{
				Error: err.Error(),
			})
			return
		}

		c.Set("info", int(userId.Id))
		c.Next()
	}
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(config.Conf.Token.CookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "bad cookie"})
			return
		}

		claims := jwthelper.Claims{}
		token, err := jwthelper.ParseToken(cookie, &claims)
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "bad token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "fake token"})
			return
		}
		c.Set("info", int(claims.Id))
		c.Next()
	}
}
