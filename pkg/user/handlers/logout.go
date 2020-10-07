package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"net/http"
	"time"
)

func Logout(c *gin.Context) {
	ss, err := c.Cookie(config.Conf.Token.CookieName)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    ss,
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, "success")
}
