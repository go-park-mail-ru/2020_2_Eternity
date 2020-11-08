package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	fstorage "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/filestorage"
	pin_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/postgres"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/usecase"
	"github.com/microcosm-cc/bluemonday"
)

func AddPinRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy, conf *config.Config) {
	rep := pin_postgres.NewRepo(db)
	store := fstorage.NewStorage(conf)
	uc := usecase.NewUsecase(rep, store)
	handler := NewHandler(uc, p)

	r.GET("/pins/board/:id", handler.GetPinsFromBoard)
	r.GET("/user/pins/:username", handler.GetAllPins)
	r.GET("/user/pin/:id", handler.GetPin)

	authorized := r.Group("/", auth.AuthCheck())

	authorized.POST("/user/pin", middleware.SendNotification(note_http.CreateNoteUsecase(db)), handler.CreatePin)


}
