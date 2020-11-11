package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	comment_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/repository/postgres"
	comment_usecase "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	"github.com/microcosm-cc/bluemonday"
)

/*
	Get   pin/comments/:comm_id	 - get concrete comment
	Get   pin/:pin_id/comments   - get all comments for pin
	Post  pin/:pin_id/comments   - create new comment for pin
*/

func AddCommentRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy) {
	rep := comment_postgres.NewRepo(db)
	uc := comment_usecase.NewUsecase(rep)
	handler := NewHandler(uc, p)

	r.POST("/pin/comments", auth.AuthCheck(), middleware.SendNotification(note_http.CreateNoteUsecase(db)), handler.CreateComment) // Use(csrf.CSRFCheck()) на фронте его еще нет, поэтому закомменчен
	r.GET(fmt.Sprintf("/pin/:%s/comments", PinIdParam), handler.GetPinComments)
	r.GET(fmt.Sprintf("/comment/:%s", CommentIdParam), handler.GetCommentById)
}
