package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/pkg/handlers"
)

type Server struct {
	config *config.ConfServer
	router *gin.Engine
}

func New(config *config.ConfServer, db *config.Database) *Server {
	r := gin.Default()

	// TODO func AddRoute
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	handler := handlers.NewHandler(db)

	r.POST("/user/signup", handler.SignUp)

	return &Server{
		config: config,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
