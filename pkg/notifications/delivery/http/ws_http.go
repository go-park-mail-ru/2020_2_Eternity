package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	usecase_ws "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/usecase"
)

type Delivery struct {
	ws *usecase_ws.WebSocketPool
}

func NewDelivery(ws *usecase_ws.WebSocketPool) *Delivery {
	return &Delivery{
		ws: ws,
	}
}

func (d * Delivery) Add(c *gin.Context) {
	//userId, ok := auth.GetClaims(c)
	//if !ok {
	//	config.Lg("ws_http", "Add").Error("Can't get claims")
	//	return
	//}

	conn, err := d.ws.Upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		config.Lg("ws_http", "Add").Error(err.Error())
		return
	}

	d.ws.AddConnection(conn, /*userId*/ 12)
}
