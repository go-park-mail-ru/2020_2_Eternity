package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/usecase"
	"github.com/microcosm-cc/bluemonday"
)

func AddBoardRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy) {
	rep := repository.NewRepo(db)
	uc := usecase.NewUsecase(rep)
	handler := NewHandler(uc, p)

	r.GET("/board/:id", handler.GetBoard)
	r.GET("/boards/:username", handler.GetAllBoardsbyUser)

	authorized := r.Group("/", auth.AuthCheck()) // authorized.Use(csrf.CSRFCheck()) на фронте его еще нет, поэтому закомменчен
	{
		authorized.POST("/board", handler.CreateBoard)
		authorized.POST("/attach", handler.AttachPinToBoard)
		authorized.DELETE("/detach", handler.DetachPinFromBoard)
	}
}
