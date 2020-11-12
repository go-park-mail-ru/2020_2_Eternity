package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
	"net/http"
)

type Delivery struct {
	uc notifications.IUsecase
}

func NewDelivery(uc notifications.IUsecase) *Delivery {
	return &Delivery{
		uc: uc,
	}
}

func (d *Delivery) GetUserNotes(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		config.Lg("notifications_http", "GetUserNotes").Error("Can't get claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	notes, err := d.uc.GetUserNotes(userId)
	if err != nil {
		config.Lg("notifications_http", "GetUserNotes").Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (d *Delivery) UpdateUserNotes(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		config.Lg("notifications_http", " UpdateUserNotes").Error("Can't get claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := d.uc.UpdateUserNotes(userId)
	if err != nil {
		config.Lg("notifications_http", " UpdateUserNotes").Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
