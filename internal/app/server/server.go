package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Server struct {
	logFile *os.File
	server  *http.Server
}

func New(config *config.Config) *Server {
	logFile := setupGinLogger()
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
		logFile: logFile,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.Web.Server.Address, config.Web.Server.Port),
			Handler: r,
		},
	}
}

func (s *Server) Run() {
	defer s.logFile.Close()

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Info("Server listening on " + s.server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

func setupGinLogger() *os.File {
	switch strings.ToLower(config.Conf.Logger.GinLevel) {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	if !config.Conf.Logger.StdoutLog {
		file, err := os.OpenFile(config.Conf.Logger.GinFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Error("Failed to log to file, using default stderr")
			return nil
		}

		gin.DefaultWriter = io.MultiWriter(file)
		return file
	} else {
		return nil
	}
}
