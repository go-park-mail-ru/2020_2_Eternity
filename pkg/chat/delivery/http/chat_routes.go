package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/microcosm-cc/bluemonday"
)

func AddChatRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy) {
	handler := NewHandler()
	r.GET("/chat", auth.AuthCheck(), handler.GetMessages)
}