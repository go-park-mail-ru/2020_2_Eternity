package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
)

func main() {
	db := database.New(database.NewConfig())
	if err := db.Open(); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
		return
	}

	srv := server.New(server.NewConfig())
	if err := srv.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
