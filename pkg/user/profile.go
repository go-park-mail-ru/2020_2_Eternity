package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"net/http"
)

func GetProfile(c *gin.Context) {
	claimsInt, ok := c.Get("info")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"can't get key"})
		return
	}

	claims, ok := claimsInt.(jwthelper.Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"can't lead claims type"})
		return
	}

	user := User{Username: claims.Username}
	if !user.GetUser() {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
