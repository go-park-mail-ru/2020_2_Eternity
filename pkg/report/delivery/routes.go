package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/report/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/report/usecase"
	"github.com/microcosm-cc/bluemonday"
)

func AddReportRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy) {
	rep := repository.New(db)
	uc := usecase.New(rep)
	handler := New(uc, p)

	authorized := r.Group("/", auth.AuthCheck()) // authorized.Use(csrf.CSRFCheck()) на фронте его еще нет, поэтому закомменчен
	{
		authorized.POST("/report", handler.ReportPin)
		authorized.GET("/reports/pin", handler.GetByPinId)
		authorized.GET("/reports/user", handler.GetByUsername)
	}
}
