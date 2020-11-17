package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/usecase"
)

func AddSearchRoute(r *gin.Engine, db database.IDbConn) {
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)

	handler := NewHandler(uc)

	r.GET("/search", handler.Search)
}
