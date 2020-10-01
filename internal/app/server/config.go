package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

func NewConfig() *Config {
	viper.SetDefault("address", "")
	viper.SetDefault("port", "8080")

	viper.SetConfigName("server_config")
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
