package search

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

func Search(c *gin.Context) {
	claimsId, ok := auth.GetClaims(c)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	sType := c.Query("type")
	cont := c.Query("content")

	switch sType {
	case "user":
		c.JSON(http.StatusOK, utils.Error{Error: fmt.Sprintf("User id = %d, wants to search by pin with content = %s", claimsId, cont)})
		return
	case "pin":
		c.JSON(http.StatusOK, utils.Error{Error: fmt.Sprintf("User id = %d, wants to search by pin with content = %s", claimsId, cont)})
		return
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: fmt.Sprintf("Bad search type %s. Can search by user and pin", sType)})
	}

}
