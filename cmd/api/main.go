package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	"log"
)

func Close() {
	if err := config.Db.Close(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	db := database.NewDB(&config.Conf.Db)
	if err := db.Open(); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	defer Close()
	srv := server.New(config.Conf, db)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
