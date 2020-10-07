package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	pinhandlers "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/handlers"
	userhandlers "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/handlers"
)

type Server struct {
	config *config.ConfServer
	router *gin.Engine
}

func New(config *config.Config) *Server {
	r := gin.Default()

	r.Static(config.Web.Static.UrlImg, config.Web.Static.DirImg)

	r.POST("/user/signup", userhandlers.SignUp)
	r.POST("/user/login", userhandlers.Login)
	r.GET("/images/avatar/:file", userhandlers.GetAvatar)

	r.POST("/user/logout", userhandlers.AuthCheck(), userhandlers.Logout)
	r.POST("/user/pin", userhandlers.AuthCheck(), pinhandlers.CreatePin)
	r.GET("/user/pin", userhandlers.AuthCheck(), pinhandlers.GetPin)
	r.GET("/user/profile", userhandlers.AuthCheck(), userhandlers.GetProfile)
	r.PUT("/user/profile/password", userhandlers.AuthCheck(), userhandlers.UpdatePassword)

	r.MaxMultipartMemory = 8 << 20
	r.POST("/user/profile/avatar", userhandlers.AuthCheck(), userhandlers.SetAvatar)

	return &Server{
		config: &config.Web.Server,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
