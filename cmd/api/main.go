package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	"log"
)

func main() {
	if err := config.Db.Open(); err != nil {
		log.Fatal(err)
		return
	}

	if err := config.Db.Close(); err != nil {
		log.Fatal(err)
		return
	}

	srv := server.New(&config.Conf.Web.Server)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
