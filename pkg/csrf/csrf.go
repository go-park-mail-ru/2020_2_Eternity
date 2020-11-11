package csrf

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
	"time"
)

func CSRFCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("X-CSRF-TOKEN")
		if value == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.Error{Error: "csrf attack"})
			return
		}
		claims := jwthelper.CsrfClaims{}
		token, err := jwthelper.ParseCsrfToken(value, &claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.Error{Error: "bad token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.Error{Error: "fake token"})
			return
		}
		t := time.Now()
		if t.After(claims.Expires) {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.Error{Error: "expires csrf"})
			return
		}
		c.Next()
	}
}
