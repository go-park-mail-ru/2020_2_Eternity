package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/usecase"
)

func AddFeedRoutes(r *gin.Engine, db database.IDbConn, conf *config.Config) {
	rep := repository.NewRepo(db)

	uc := usecase.NewUseCase(rep)
	handler := NewHandler(uc)

	r.GET("/feed", handler.GetFeed)
	r.Use(auth.AuthCheck())
	r.GET("/subfeed", handler.GetSubFeed)
}
