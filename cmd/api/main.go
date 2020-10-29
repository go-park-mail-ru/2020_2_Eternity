package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	log "github.com/sirupsen/logrus"
)

func init() {
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()
}

func main() {
	logger := config.Logger{}
	logger.Init()
	defer logger.Cleanup()

	defer Close()

	if conn := config.Db; conn == nil {
		config.Lg("main", "main").Fatal("Connection refused")
		return
	}
	config.Lg("main", "main").Info("Connected to DB")

	srv := server.New(config.Conf)

	srv.Run()

	config.Lg("main", "main").Info("Server stopped")
}

func Close() {
	if err := config.Db.Close(); err != nil {
		log.Fatal(err)
		return
	}
	config.Lg("main", "Close").Info("DB connection closed")
}
