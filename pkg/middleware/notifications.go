package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
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

		if err := uc.CreateNotes(note); err != nil {
			config.Lg("Middleware", "SendNotification").Error(err.Error())
			return
		}

	}
}
