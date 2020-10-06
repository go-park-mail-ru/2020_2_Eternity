package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/handlers"
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

	r.POST("/user/signup", handlers.SignUp)
	r.POST("/user/login", handlers.Login)

	r.POST("/user/logout", handlers.AuthCheck(), handlers.Logout)
	r.PUT("/user/profile", handlers.AuthCheck(), handlers.UpdateUser)
	r.PUT("/user/profile/password", handlers.AuthCheck(), handlers.UpdatePassword)

	r.MaxMultipartMemory = 8 << 20
	r.POST("/user/profile/avatar", handlers.AuthCheck(), handlers.SetAvatar)
	r.GET("/user/profile/avatar", handlers.AuthCheck(), handlers.GetAvatar)
	return &Server{
		config: config,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
