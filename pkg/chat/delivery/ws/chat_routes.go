package ws

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)

func AddChatWsRoutes(server ws.IServer, conn grpc.ClientConnInterface,) {
	uc := usecase.NewUsecase(conn)
	handler := NewHandler(uc, bluemonday.UGCPolicy())

	server.SetHandler("get_ass", handler.CreateMessage)
}
