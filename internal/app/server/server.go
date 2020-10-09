package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
)

type Server struct {
	config *config.ConfServer
	router *gin.Engine
}

func New(config *config.Config) *Server {
	r := gin.Default()

	r.Static(config.Web.Static.UrlImg, config.Web.Static.DirImg)

	r.POST("/user/signup", user.SignUp)
	r.POST("/user/login", user.Login)
	r.GET("/images/avatar/:file", user.GetAvatar)

	r.POST("/user/logout", user.AuthCheck(), user.Logout)
	r.POST("/user/pin", user.AuthCheck(), pin.CreatePin)
	r.GET("/user/pin", user.AuthCheck(), pin.GetPin)
	r.GET("/user/profile", user.AuthCheck(), user.GetProfile)

	r.PUT("/user/profile/password", user.AuthCheck(), user.UpdatePassword)
	r.PUT("/user/profile/", user.AuthCheck(), user.UpdateUser)

	r.MaxMultipartMemory = 8 << 20
	r.POST("/user/profile/avatar", user.AuthCheck(), user.SetAvatar)

	return &Server{
		config: &config.Web.Server,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
