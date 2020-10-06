package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
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

	user := model.User{Username: claims.Username}
	if !user.GetUser() {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"user not found"})
		return
	}

	profile := api.GetProfileApi{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}

	c.JSON(http.StatusOK, profile)
}
