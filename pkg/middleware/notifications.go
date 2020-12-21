package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
)

func SendNotification(uc notifications.IUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		note, ok := c.Get(domain.NotificationKey)
		if !ok {
			config.Lg("Middleware", "SendNotification").
				Error("Notification not given")
			return
		}

		if err := uc.CreateNotes(&note); err != nil {
			config.Lg("Middleware", "SendNotification").Error(err.Error())
			return
		}

	}
}

func SendNotificationWs(uc notifications.IUsecase) func(c *ws.Context) {
	return func(c *ws.Context) {
		c.Next()

		note, err := c.Get(domain.NotificationKey)
		if err != nil {
			config.Lg("Middleware", "SendNotificationWs").
				Error("Notification not given: ", err.Error())
			return
		}

		if err := uc.CreateNotes(&note); err != nil {
			config.Lg("Middleware", "SendNotificationWs").Error(err.Error())
			return
		}

	}
}