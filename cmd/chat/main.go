package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	chatServer "github.com/go-park-mail-ru/2020_2_Eternity/internal/chat"
)

func main() {
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()

	logger := config.Logger{}
	logger.Init()
	defer func() {
		if err := logger.Cleanup(); err != nil {
			config.Lg("chat", "main").Error(err.Error())
		}
	}()

	dbConn := database.NewDB(&config.Conf.Db)
	if err := dbConn.Open(); err != nil {
		config.Lg("main", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("main", "main").Info("Connected to DB")

	srv := chatServer.New(dbConn)
	srv.Run()

	config.Lg("main", "main").Info("Server stopped")
}
