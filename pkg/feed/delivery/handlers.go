package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
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
	pins, err := h.uc.GetFeed(0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Error{Error: "Cant get feed"})
		return
	}

	c.JSON(http.StatusOK, pins)
}
