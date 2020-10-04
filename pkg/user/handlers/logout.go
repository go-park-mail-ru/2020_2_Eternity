package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Logout(c *gin.Context) {
	ss, err := c.Cookie(cookiename)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     cookiename,
		Value:    ss,
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, "success")
}