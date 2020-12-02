package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/middleware"
	note_http "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	proto_auth "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)


func AddUserRoutes(r *gin.Engine, db database.IDbConn, p *bluemonday.Policy, ac grpc.ClientConnInterface,  server ws.IServer) {
	rep := repository.NewRepo(db)
	uc := usecase.NewUsecase(rep)
	client := proto_auth.NewAuthServiceClient(ac)
	handler := NewHandler(uc, p, client)

	mw := auth.NewAuthMw(client)

	r.POST("/user/signup", handler.SignUp)
	r.POST("/user/login", handler.Login)
	r.GET("/images/avatar/:file", handler.GetAvatar)
	r.GET("/followers/:username", handler.GetFollowers)
	r.GET("/following/:username", handler.GetFollowing)
	r.GET("/userpage/:username", handler.GetUserPage)
	authorized := r.Group("/")
	authorized.Use(mw.AuthCheck()) // authorized.Use(csrf.CSRFCheck()) на фронте его еще нет, поэтому закомменчен
	{
		authorized.POST("/user/logout", handler.Logout)
		authorized.GET("/user/profile", handler.GetProfile)
		authorized.PUT("/user/profile/password", handler.UpdatePassword)
		authorized.PUT("/user/profile", handler.UpdateUser)
		authorized.POST("/user/profile/avatar", handler.SetAvatar)

		authorized.Group("/", middleware.SendNotification(note_http.CreateNoteUsecase(db, server))).
			POST("/follow", handler.Follow)

		authorized.POST("/unfollow", handler.Unfollow)
		authorized.GET("/isfollowing/:username", handler.IsFollowing)
	}
}
