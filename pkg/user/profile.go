package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

func GetProfile(c *gin.Context) {
	claimsId, ok := auth.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Error{Error: "can't get key"})
		return
	}

	user := User{ID: claimsId}
	if !user.GetUser() {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
