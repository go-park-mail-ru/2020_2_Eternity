package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"net/http"
)



type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetMessages(c *gin.Context) {
	_, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("comment_http", "CreateComment").Error("Can't get claims")
		return
	}

	c.JSON(http.StatusOK, "{}")
}


