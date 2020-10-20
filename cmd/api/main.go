package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	log "github.com/sirupsen/logrus"
)

func Close() {
	if err := config.Db.Close(); err != nil {
		log.Fatal(err)
		return
	}
	log.Info("DB connection closed")
}

func main() {
	logger := config.Logger{}
	logger.Init()
	defer logger.Cleanup()

	defer Close()

	log.Info("ConfLogger started")
	log.WithFields(log.Fields{"jopa": "benis"}).Info("Benis")

	if conn := config.Db; conn == nil {
		log.Fatal("Connection refused")
		return
	}
	log.Info("Connected to DB")

	srv := server.New(config.Conf)

	srv.Run()

	log.Info("Server exiting")
}
