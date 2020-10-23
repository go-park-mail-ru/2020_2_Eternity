package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"net/http"
)

func GetPin(c *gin.Context) {
	claimsId, ok := user.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("pin", "GetPin").Error("Can't get claims")
		return
	}

	pins, err := GetPinList(claimsId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin", "GetPin").Error("GetPinList: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, pins)
}
