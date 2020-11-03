package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets/usecase"
)

func WsMiddleware(ws *usecase.WebSocketPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		note, ok := GetNotefication(c)
		if !ok {
			config.Lg("ws_moddleware", "WsMiddleware").Error("Can't get notification")
			return
		}

		ws.AddNotification(&note)
		config.Lg("ws_moddleware", "WsMiddleware").Info("Got note ", note)
	}
}


func GetNotefication(c *gin.Context) (domain.Notification, bool) {
	claims, ok := c.Get("note")
	if !ok {
		return domain.Notification{}, false
	}

	claimsId, ok := claims.(domain.Notification)
	if !ok {
		return domain.Notification{}, false
	}
	return claimsId, true
}