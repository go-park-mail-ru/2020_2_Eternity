package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"net/http"
)

func GetPin(c *gin.Context) {
	claims, ok := c.Get("info")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, user.Error{"can't get key"})
		return
	}

	requester, ok := claims.(jwthelper.Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"can't lead claims"})
		return
	}

	pins, err := GetPinList(requester.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[GetPinList]: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, pins)
}
