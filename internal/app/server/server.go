package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *Config
	router *gin.Engine
}

func New(config *Config) *Server {
	r := gin.Default()

	// TODO func AddRoute
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return &Server{
		config: config,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}
