package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	boardDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/delivery"
	feedDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/delivery"
	pinDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/delivery/http"
	userDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/delivery"
	"github.com/microcosm-cc/bluemonday"

	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets"
)

type Server struct {
	logFile *os.File
	server  *http.Server
}

func New(config *config.Config, db database.IDbConn) *Server {
	logFile := setupGinLogger()

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Static(config.Web.Static.UrlImg, config.Web.Static.DirImg)

	ws := websockets.NewPool()
	r.GET("/ws", ws.Add)

	p := bluemonday.UGCPolicy()

	userDelivery.AddUserRoutes(r, db, p, ws)
	pinDelivery.AddPinRoutes(r, db, p, config)
	boardDelivery.AddBoardRoutes(r, db, p)

	feedDelivery.AddFeedRoutes(r, db, config)

	rpd := comment.NewResponder()
	r.POST("/pin/comments", auth.AuthCheck(), rpd.CreateComment)
	r.GET(fmt.Sprintf("/pin/:%s/comments", comment.PinIdParam), rpd.GetComments)
	r.GET(fmt.Sprintf("/comment/:%s", comment.CommentIdParam), rpd.GetCommentById)

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
	config.Lg("server", "Run").Info("Server listening on " + s.server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	config.Lg("server", "Run").Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		config.Lg("server", "Run").Fatal("Server forced to shutdown:", err)
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
			config.Lg("server", "setupGinLogger").Fatal("Failed to log to file, using default stderr")
			return nil
		}

		gin.DefaultWriter = io.MultiWriter(file)
		return file
	} else {
		return nil
	}
}
