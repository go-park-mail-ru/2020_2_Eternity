package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

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
		c.Set("info", claims.Id)
		c.Next()
	}
}

func GetClaims(c *gin.Context) (int, bool) {
	claims, ok := c.Get("info")
	if !ok {
		return -1, false
	}

	claimsId, ok := claims.(int)
	if !ok {
		return -1, false
	}
	return claimsId, true
}
