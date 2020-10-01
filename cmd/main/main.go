package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/handler"
)

func main() {
	dbConfig, err := configs.ReadConfig()
	if err != nil {
		return
	}

	db, err := database.InitDB(dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	signupHandler := handler.NewHandler(db)

	r := gin.Default()
	r.POST("/user/signup", signupHandler.SignUp)
	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}
