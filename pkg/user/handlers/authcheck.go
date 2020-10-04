package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(config.Conf.Token.CookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"bad cookie"})
			return
		}
		claims := jwthelper.Claims{}
		token, err := jwthelper.ParseToken(cookie, &claims)
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"bad token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"fake token"})
			return
		}
		c.Set("info", claims)
		c.Next()
	}
}
