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

func New(config *config.ConfServer) *Server {
	r := gin.Default()

	// TODO func AddRoute
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/user/signup", userhandlers.SignUp)
	r.POST("/user/login", userhandlers.Login)
	r.POST("/user/logout", userhandlers.AuthCheck(), userhandlers.Logout)
	r.POST("/user/pin", userhandlers.AuthCheck(), pinhandlers.CreatePin) // CreatePin

	return &Server{
		config: config,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
