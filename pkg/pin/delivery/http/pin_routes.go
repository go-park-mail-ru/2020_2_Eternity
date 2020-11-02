package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	fstorage "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/filestorage"
	pin_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/postgres"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/usecase"
)

func AddPinRoutes(r *gin.Engine, db database.IDbConn, conf *config.Config) {
	rep := pin_postgres.NewRepo(db)
	store := fstorage.NewStorage(conf)
	uc := usecase.NewUsecase(rep, store)
	handler := NewHandler(uc)

	authorized := r.Group("/", auth.AuthCheck())

	authorized.POST("/user/pin", handler.CreatePin)
	authorized.GET("/user/pin", handler.GetAllPins)
}
