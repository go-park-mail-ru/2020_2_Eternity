package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	grpcAuth "github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth/delivery/grpc"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/usecase"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()

	logger := config.Logger{}
	logger.Init()
	defer logger.Cleanup()

	dbConn := database.NewDB(&config.Conf.Db)
	if err := dbConn.Open(); err != nil {
		config.Lg("authserv", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("authserv", "main").Info("Connected to DB")

	repo := repository.NewRepo(dbConn)
	uc := usecase.NewUsecase(repo)
	handler := grpcAuth.NewHandler(uc)

	lis, err := net.Listen(config.Conf.Web.Search.Protocol, config.Conf.Web.Auth.Address+":"+config.Conf.Web.Auth.Port)
	if err != nil {
		config.Lg("searchserv", "main").Fatal(err.Error())
	}

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, handler)

	if err := server.Serve(lis); err != nil {
		config.Lg("authserv", "main").Fatal("cant run server")
	}
	config.Lg("authserv", "main").Info("server started")
}
