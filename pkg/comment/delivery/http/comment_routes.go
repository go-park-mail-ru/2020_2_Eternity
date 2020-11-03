package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	comment_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/repository/postgres"
	comment_usecase "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/usecase"
	ws_middleware "github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets/middleware"
	ws_usecase "github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets/usecase"
)

/*
	Get   pin/comments/:comm_id	 - get concrete comment
	Get   pin/:pin_id/comments   - get all comments for pin
	Post  pin/:pin_id/comments   - create new comment for pin
*/


func AddCommentRoutes(r *gin.Engine, db database.IDbConn, ws *ws_usecase.WebSocketPool) {
	rep := comment_postgres.NewRepo(db)
	uc := comment_usecase.NewUsecase(rep)
	handler := NewHandler(uc)

	r.POST("/pin/comments", auth.AuthCheck(), ws_middleware.WsMiddleware(ws), handler.CreateComment)
	r.GET(fmt.Sprintf("/pin/:%s/comments", PinIdParam), handler.GetPinComments)
	r.GET(fmt.Sprintf("/comment/:%s", CommentIdParam), handler.GetCommentById)
}
