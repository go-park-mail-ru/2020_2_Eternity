package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	ws2 "github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"google.golang.org/grpc"
	"os"
)

func initDirs() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(root+config.Conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		config.Lg("create dirs", "initDirs").
			Error("MkAllDir: ", err.Error())
		return err
	}
	if err := os.MkdirAll(root+config.Conf.Web.Static.DirAvt, 0777|os.ModeDir); err != nil {
		config.Lg("create dirs", "initDirs").
			Error("MkAllDir: ", err.Error())
		return err
	}
	return nil
}

func InitClientConnections() (*grpc.ClientConn, *grpc.ClientConn, error) {
	sc, err := grpc.Dial(config.Conf.Web.Search.Address+":"+config.Conf.Web.Search.Port,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	ac, err := grpc.Dial(config.Conf.Web.Auth.Address+":"+config.Conf.Web.Auth.Port,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	return sc, ac, nil
}

func main() {
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()

	logger := config.Logger{}
	logger.Init()
	defer func() {
		if err := logger.Cleanup(); err != nil {
			config.Lg("main", "main").Error(err.Error())
		}
	}()

	defer config.Db.Close() // NOTE (Pavel S) Temporary

	if err := initDirs(); err != nil {
		config.Lg("main", "main").Fatal("Cannot create dirs")
		return
	}

	dbConn := database.NewDB(&config.Conf.Db)
	if err := dbConn.Open(); err != nil {
		config.Lg("main", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("main", "main").Info("Connected to DB")

	sc, ac, err := InitClientConnections()
	if err != nil {
		config.Lg("main", "search").Fatal("Connection grpc refused")
		return
	}
	defer sc.Close()
	defer ac.Close()

	wsSrv := ws2.NewServer()
	wsSrv.Run()
	defer wsSrv.Stop()

	chMsConn := server.NewChatMsConnection()
	defer chMsConn.Close()

	srv := server.New(config.Conf, dbConn, sc, ac, chMsConn, wsSrv)
	srv.Run()

	config.Lg("main", "main").Info("Server stopped")
}
