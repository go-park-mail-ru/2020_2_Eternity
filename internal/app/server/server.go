package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	chatDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/delivery/http"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat/delivery/ws"
	commentDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/delivery/http"
	search "github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/delivery/http"

	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/metric"

	ws2 "github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"

	noteDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/delivery/http"
	"google.golang.org/grpc"

	boardDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/delivery"
	feedDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/delivery"
	pinDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/delivery/http"
	reportDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/report/delivery"
	userDelivery "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/delivery"
	"github.com/microcosm-cc/bluemonday"

	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
)

type Server struct {
	logFile *os.File
	server  *http.Server
}

func New(conf *config.Config, db database.IDbConn, sc *grpc.ClientConn, ac *grpc.ClientConn,
	chMsConn grpc.ClientConnInterface, wsSrv ws2.IServer) *Server {

	logFile := setupGinLogger()

	r := gin.Default()

	r.RouterGroup = *r.Group("/api")

	r.MaxMultipartMemory = 8 << 20
	r.Static(conf.Web.Static.UrlImg, conf.Web.Static.DirImg)

	m, err := metric.CreateNewMetric("main")

	if err != nil {
		log.Fatal("errr", err)
	}

	r.Use(metric.CollectMetrics(m))

	go metric.RouterForMetrics(conf.Monitoring.Api.Address + ":" + conf.Monitoring.Api.Port)

	noteDelivery.AddNoteRoutes(r, db, wsSrv)

	p := bluemonday.UGCPolicy()

	search.AddSearchRoute(r, sc)
	chatDelivery.AddChatRoutes(r, db, chMsConn, p, wsSrv)
	ws.AddChatWsRoutes(wsSrv, db, chMsConn)
	commentDelivery.AddCommentRoutes(r, db, p, wsSrv)
	userDelivery.AddUserRoutes(r, db, p, ac, wsSrv)
	pinDelivery.AddPinRoutes(r, db, p, conf, wsSrv)
	boardDelivery.AddBoardRoutes(r, db, p)
	reportDelivery.AddReportRoutes(r, db, p)
	feedDelivery.AddFeedRoutes(r, db, conf)

	return &Server{
		logFile: logFile,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", conf.Web.Server.Address, conf.Web.Server.Port),
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

	quit := make(chan os.Signal, 1)
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
