package search

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
)

func AddSearchRoute(r *gin.Engine) {
	r.GET("/search", auth.AuthCheck(), Search)
}
