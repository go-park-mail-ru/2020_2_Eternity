package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/usecase"
)

func AddUserRoutes(r *gin.Engine, db database.IDbConn) {
	rep := repository.NewRepo(db)
	uc := usecase.NewUsecase(rep)
	handler := NewHandler(uc)

	r.POST("/user/signup", handler.SignUp)
	r.POST("/user/login", handler.Login)
	r.GET("/images/avatar/:file", handler.GetAvatar)

	authorized := r.Group("/")
	authorized.Use(auth.AuthCheck())
	{
		authorized.POST("/user/logout", handler.Logout)
		authorized.GET("/user/profile", handler.GetProfile)
		authorized.PUT("/user/profile/password", handler.UpdatePassword)
		authorized.PUT("/user/profile/", handler.UpdateUser)
		authorized.POST("/user/profile/avatar", handler.SetAvatar)

		authorized.Group("/", middleware.SendNotification(note_http.CreateNoteUsecase(db))).
			POST("/follow", handler.Follow)

		authorized.POST("/unfollow", handler.Unfollow)
	}
}
