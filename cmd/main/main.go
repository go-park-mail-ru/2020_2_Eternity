package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/handler"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
)



func main() {
	users := model.NewMockUsers()
	signupHandler := handler.NewHandler(users)

	r := gin.Default()
	r.POST("/user/signup", signupHandler.SignUp)
	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}