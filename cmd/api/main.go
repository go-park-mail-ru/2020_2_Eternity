package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
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
	if conn := config.Db; conn == nil {
		log.Fatal("Connection refused")
		return
	}

	defer Close()
	srv := server.New(config.Conf)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
