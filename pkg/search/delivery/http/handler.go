package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	uc search.IUsecase
}

func NewHandler(uc search.IUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Search(c *gin.Context) {
	/*claimsId, ok := jwthelper.GetClaims(c)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}*/

	sType := c.Query("type")
	cont := c.Query("content")
	if cont == "" {
		c.JSON(http.StatusBadRequest, utils.Error{Error: "no content"})
		return
	}
	last, err := strconv.Atoi(c.Query("last"))
	if err != nil {
		last = 0
	}

	switch sType {
	case "user":
		users, err := h.uc.GetUsersByName(cont, last)
		if err != nil {
			c.JSON(http.StatusOK, utils.Error{Error: "not found"})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	case "pin":
		pins, err := h.uc.GetPinsByTitle(cont, last)
		if err != nil {
			c.JSON(http.StatusOK, utils.Error{Error: "pins not found"})
			return
		}
		c.JSON(http.StatusOK, pins)
		return
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: fmt.Sprintf("Bad search type %s. Can search by user and pin", sType)})
	}

}
