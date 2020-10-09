package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(c *gin.Context) {
	claims, ok := isAuthorised(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"can't get key"})
		return
	}

	user := User{Username: claims.Username}
	if !user.GetUser() {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
