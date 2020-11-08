package main

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/server"
	"os"
)

func init() {
	// NOTE (Pavel S) Temporary
	config.Conf = config.NewConfig()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()
}

func initDirs() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(root+config.Conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		config.Lg("create dirs", "initDirs").
			Error("MkAllDir: ", err.Error())
		return err
	}
	if err := os.MkdirAll(root+config.Conf.Web.Static.DirAvt, 0777|os.ModeDir); err != nil {
		config.Lg("create dirs", "initDirs").
			Error("MkAllDir: ", err.Error())
		return err
	}
	return nil
}

func main() {
	logger := config.Logger{}
	logger.Init()
	defer logger.Cleanup()

	defer config.Db.Close() // NOTE (Pavel S) Temporary

	if err := initDirs(); err != nil {
		config.Lg("main", "main").Fatal("Cannot create dirs")
		return
	}

	dbConn := database.NewDB(&config.Conf.Db)
	if err := dbConn.Open(); err != nil {
		config.Lg("main", "main").Fatal("Connection refused")
		return
	}
	defer dbConn.Close()
	config.Lg("main", "main").Info("Connected to DB")

	srv := server.New(config.Conf, dbConn)

	srv.Run()

	config.Lg("main", "main").Info("Server stopped")
}
