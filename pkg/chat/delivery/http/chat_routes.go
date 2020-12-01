package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	ws2 "github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)

func AddChatRoutes(r *gin.Engine, db database.IDbConn, conn grpc.ClientConnInterface, p *bluemonday.Policy, srv ws2.IServer) {
	uc := usecase.NewUsecase(conn)
	handler := NewHandler(uc, p)

	authorized := r.Group("/", auth.AuthCheck())

	authorized.GET("/ws", ServeWs(srv))
	authorized.POST("/chat" , middleware.SendNotification(note_http.CreateNoteUsecase(db, srv)), handler.CreateChat)
	authorized.GET("/chat/:" + ChatIdParam, handler.GetChatById)
	authorized.GET("/chat", handler.GetUserChats)
	authorized.PUT("/chat", handler.MarkAllMessagesRead)
}
