package chat

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	chatGrpc "github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat/delivery/grpc"
	chatRepo "github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat/usecase"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/chat"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	srv *grpc.Server
	lis net.Listener
}

func New(db database.IDbConn) *Server {
	l, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%s", config.Conf.Web.Chat.Address, config.Conf.Web.Chat.Port),
	)

	if err != nil {
		config.Lg("chat_server", "New").Fatal("Failed to listen: ", err.Error())
	}


	repo := chatRepo.NewRepo(db)
	uc := usecase.NewUsecase(repo)
	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, chatGrpc.NewChatServer(uc))


	return &Server{
		srv: grpcServer,
		lis: l,
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.srv.Serve(s.lis); err != nil {
			config.Lg("chat_server", "Run").Fatal("Can't serve: ", err.Error())
			return
		}
	}()

	config.Lg("chat_server", "Run").Info("Server listening on " +
		fmt.Sprintf("%s:%s", config.Conf.Web.Chat.Address, config.Conf.Web.Chat.Port))

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	config.Lg("chat_server", "Run").Info("Shutting down server...")

	s.srv.Stop()
}
