package database

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DriverName string `mapstructure:"driver_name"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DbName     string `mapstructure:"db_name"`
	SslMode    string `mapstructure:"ssl_mode"`
}


func NewConfig() *Config {
	viper.SetDefault("driver_name", "default")
	viper.SetDefault("username", "default")
	viper.SetDefault("password", "0000")
	viper.SetDefault("db_name", "default")
	viper.SetDefault("ssl_mode", "disable")

	viper.SetConfigName("database_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	conf := new(Config)

	er := viper.Unmarshal(conf)
	if er != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return conf
}
