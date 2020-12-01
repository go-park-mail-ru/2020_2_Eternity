package ws

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/usecase"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)

func AddChatWsRoutes(server ws.IServer, db database.IDbConn, conn grpc.ClientConnInterface) {
	uc := usecase.NewUsecase(conn)
	handler := NewHandler(uc, bluemonday.UGCPolicy())

	server.SetHandler(
		domainChat.CreateMessageReqType,
		middleware.SendNotificationWs(note_http.CreateNoteUsecase(db, server)),
		handler.CreateMessage,
	)

	server.SetHandler(domainChat.DeleteMessageReqType, handler.DeleteMessage)
	server.SetHandler(domainChat.GetLastNMessagesReqType, handler.GetLastNMessages)
	server.SetHandler(domainChat.GetNMessagesBeforeReqType, handler.GetNMessagesBefore)
}
