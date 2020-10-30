package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/delivery"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets"
	"log"
)

type Server struct {
	config *config.ConfServer
	router *gin.Engine
}

func TestMwWs(ws *websockets.WebSocketPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status, ok := c.Get("status")
		if !ok {
			log.Println("STATUS")
			return
		}
		uid, ok := c.Get("follow_id")
		if !ok {
			log.Println("ID IS NOT SET")
			return
		}
		ws.Send(uid.(int), []byte("kto-to podpisalsya grats!"))
		log.Println(status, uid)
	}
}

func New(config *config.Config, db database.DBInterface) *Server {

	ws := websockets.NewPool()

	rep := repository.NewRepo(db)
	uc := usecase.NewUsecase(rep)

	handler := delivery.NewHandler(uc)

	r := gin.Default()

	r.Static(config.Web.Static.UrlImg, config.Web.Static.DirImg)

	r.POST("/user/signup", handler.SignUp)
	r.POST("/user/login", handler.Login)
	r.GET("/images/avatar/:file", handler.GetAvatar)
	//r.GET("/profile/:username", handler.GetUserPage)
	r.GET("/ws", ws.Add)

	r.Use(auth.AuthCheck())

	r.POST("/user/logout", handler.Logout)
	r.POST("/user/pin", pin.CreatePin)
	r.GET("/user/pin", pin.GetPin)
	r.GET("/user/profile", handler.GetProfile)

	r.PUT("/user/profile/password", handler.UpdatePassword)
	r.PUT("/user/profile/", handler.UpdateUser)

	r.MaxMultipartMemory = 8 << 20
	r.POST("/user/profile/avatar", handler.SetAvatar)
	// experimental
	r.POST("/follow", TestMwWs(ws), handler.Follow)
	r.POST("/unfollow", handler.Unfollow)
	return &Server{
		config: &config.Web.Server,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
