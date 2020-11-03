package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/websockets/usecase"
)

func init() {
	// NOTE (Pavel S) Temporary
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()
}

func main() {
	logger := config.Logger{}
	logger.Init()
	defer logger.Cleanup()

	defer config.Db.Close() // NOTE (Pavel S) Temporary

	dbConn := database.NewDB(&config.Conf.Db)
	if err := dbConn.Open(); err != nil {
		config.Lg("main", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("main", "main").Info("Connected to DB")

	ws := usecase.NewPool()
	ws.Run()
	defer ws.Stop()

	srv := server.New(config.Conf, dbConn, ws)

	srv.Run()

	config.Lg("main", "main").Info("Server stopped")
}
