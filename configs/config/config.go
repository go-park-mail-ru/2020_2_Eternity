package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Db  ConfDB  `mapstructure:"database"`
	Web ConfWeb `mapstructure:"web"`
}

type ConfDB struct {
	Postgres ConfPostgres `mapstructure:"postgres"`
}

type ConfPostgres struct {
	DriverName string `mapstructure:"driver_name"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DbName     string `mapstructure:"db_name"`
	SslMode    string `mapstructure:"ssl_mode"`
	Host       string `mapstructure:"host"`
}

type ConfWeb struct {
	Server ConfServer `mapstructure:"server"`
}

type ConfServer struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

func NewConfig() *Config {
	setDefaultDb()
	setDefaultWeb()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	conf := new(Config)

	er := viper.Unmarshal(conf)
	if er != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	fmt.Println(conf)

	return conf
}

func setDefaultDb() {
	viper.SetDefault("database.postgres.driver_name", "default")
	viper.SetDefault("database.postgres.username", "default")
	viper.SetDefault("database.postgres.password", "0000")
	viper.SetDefault("database.postgres.db_name", "default")
	viper.SetDefault("database.postgres.ssl_mode", "disable")
	viper.SetDefault("database.postgres.host", "disable")
}

func setDefaultWeb() {
	viper.SetDefault("web.server.address", "")
	viper.SetDefault("web.server.port", "0000")
}
