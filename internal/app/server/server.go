package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
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

func New(config *config.Config) *Server {

	ws := websockets.NewPool()

	r := gin.Default()

	r.Static(config.Web.Static.UrlImg, config.Web.Static.DirImg)

	r.POST("/user/signup", user.SignUp)
	r.POST("/user/login", user.Login)
	r.GET("/images/avatar/:file", user.GetAvatar)
	r.GET("/profile/:username", user.GetUserPage)
	r.GET("/ws", ws.Add)

	r.Use(auth.AuthCheck())

	r.POST("/user/logout", user.Logout)
	r.POST("/user/pin", pin.CreatePin)
	r.GET("/user/pin", pin.GetPin)
	r.GET("/user/profile", user.GetProfile)

	r.PUT("/user/profile/password", user.UpdatePassword)
	r.PUT("/user/profile/", user.UpdateUser)

	r.MaxMultipartMemory = 8 << 20
	r.POST("/user/profile/avatar", user.SetAvatar)
	// experimental
	r.POST("/follow", TestMwWs(ws), user.Follow)
	r.POST("/unfollow", user.Unfollow)
	return &Server{
		config: &config.Web.Server,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
