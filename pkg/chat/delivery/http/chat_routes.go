package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/usecase"
	ws2 "github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)

func AddChatRoutes(r *gin.Engine, conn grpc.ClientConnInterface, p *bluemonday.Policy, srv ws2.IServer) {
	uc := usecase.NewUsecase(conn)
	handler := NewHandler(uc, p)

	r.GET("/chat/ws" /*auth.AuthCheck(),*/, ServeWs(srv))
	r.POST("/chat", /*auth.AuthCheck(),*/ handler.CreateChat)
}
