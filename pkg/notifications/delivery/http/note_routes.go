package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
	note_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/repository/postgres"
	chat_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/usecase"
	pin_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/postgres"
	user_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
)

func AddNoteRoutes(r *gin.Engine, db database.IDbConn, server ws.IServer) {
	delivery := NewDelivery(CreateNoteUsecase(db, server))

	authorized := r.Group("/", auth.AuthCheck()) // authorized.Use(csrf.CSRFCheck()) на фронте его еще нет, поэтому закомменчен

	authorized.GET("/notifications", delivery.GetUserNotes)
	authorized.PUT("/notifications", delivery.UpdateUserNotes)
}

func CreateNoteUsecase(db database.IDbConn, server ws.IServer) notifications.IUsecase {
	repPin := pin_postgres.NewRepo(db)
	repUser := user_postgres.NewRepo(db)
	repNote := note_postgres.NewRepo(db)
	repChat := chat_postgres.NewRepo(db)
	return usecase.NewUsecase(repNote, repPin, repUser, repChat, server)
}
