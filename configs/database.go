package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func ReadConfig() (*DBConfig, error) {
	viper.SetConfigName("database")
	viper.AddConfigPath("configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &DBConfig{
		Host:     viper.Get("database.host").(string),
		Port:     uint16(viper.Get("database.port").(int)),
		Database: viper.Get("database.dbname").(string),
		User:     viper.Get("database.user").(string),
		Password: viper.Get("database.password").(string),
	}, nil
}
