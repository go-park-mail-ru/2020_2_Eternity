package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

func GetUserPage(c *gin.Context) {
	u := User{
		Username: c.Param("username"),
	}
	if err := u.GetUserByNameWithFollowers(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "not found user"})
		return
	}

	pins, err := pin.GetPinList(u.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.Error{Error: "pins not found"})
		return
	}
	userPage := api.UserPage{
		Username:  u.Username,
		Followers: u.Followers,
		Following: u.Following,
		PinsList:  pins,
	}
	c.JSON(http.StatusOK, userPage)
}
