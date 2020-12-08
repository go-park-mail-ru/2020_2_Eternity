package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	uc feed.IUseCase
}

func NewHandler(uc feed.IUseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) GetFeed(c *gin.Context) {
	last := c.Query("last")
	id, err := strconv.Atoi(last)
	if err != nil {
		id = 0
	}
	pins, err := h.uc.GetFeed(0, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Error{Error: "Cant get feed"})
		return
	}

	c.JSON(http.StatusOK, pins)
}

func (h *Handler) GetSubFeed(c *gin.Context) {

	last := c.Query("last")
	lastId, err := strconv.Atoi(last)
	if err != nil {
		lastId = 0
	}

	claimsId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}
	pins, err := h.uc.GetSubFeed(claimsId, lastId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Error{Error: "Cant get feed"})
		return
	}

	c.JSON(http.StatusOK, pins)
}
