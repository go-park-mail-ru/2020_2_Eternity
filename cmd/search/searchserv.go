package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/metric"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	grpcSearch "github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/delivery/grpc"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/usecase"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"log"
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
		config.Lg("searchserv", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("searchserv", "main").Info("Connected to DB")

	go metric.RouterForMetrics("localhost:7008")

	m, err := metric.CreateNewMetric("search")
	interceptor := metric.NewInterceptor(m)

	if err != nil {
		log.Fatal(err)
		return
	}

	repo := repository.NewRepository(dbConn)
	uc := usecase.NewUsecase(repo)
	handler := grpcSearch.NewHandler(uc)

	lis, err := net.Listen(config.Conf.Web.Search.Protocol, config.Conf.Web.Search.Address+":"+config.Conf.Web.Search.Port)
	if err != nil {
		config.Lg("searchserv", "main").Fatal(err.Error())
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpc_recovery.UnaryServerInterceptor(), interceptor.Collect),
	)
	search.RegisterSearchServiceServer(server, handler)

	if err := server.Serve(lis); err != nil {
		config.Lg("searchserv", "main").Fatal("cant run server")
	}
	config.Lg("searchserv", "main").Info("server started")
}
