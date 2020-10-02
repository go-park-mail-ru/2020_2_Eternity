package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	"log"
)

func main() {
	conf := config.NewConfig()

	db := database.New(&conf.Db)
	if err := db.Open(); err != nil {
		log.Fatal(err)
		return
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
		return
	}

	srv := server.New(&conf.Web.Server)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
