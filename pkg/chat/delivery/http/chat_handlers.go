package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/delivery/ws"
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

func ServeWs(h *ws.Hub) func(c *gin.Context) {
	return func(c *gin.Context) {
		//userId, ok := jwthelper.GetClaims(c)
		//if !ok {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	config.Lg("pin_http", "ServeWs").Error("Can't get claims")
		//	return
		//}

		userId := 2 // Note: for tests

		conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			config.Lg("chat_http", "ServeWs").Error(err.Error())
			return
		}

		client := ws.NewClient(h, conn, userId)
		client.Register()
	}
}
