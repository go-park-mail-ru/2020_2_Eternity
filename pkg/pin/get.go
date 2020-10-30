package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"log"
	"net/http"
)

func GetPin(c *gin.Context) {
	claimsId, ok := auth.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Println("[GetPin]: Can't get claims")
		return
	}

	pins, err := GetPinList(claimsId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("[GetPin]-[GetPinList]: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, pins)
}
