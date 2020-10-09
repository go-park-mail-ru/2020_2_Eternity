package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Conf = newConfig()
	Db   = newDatabase(&Conf.Db).Open()
)

type Config struct {
	Db    ConfDB    `mapstructure:"database"`
	Web   ConfWeb   `mapstructure:"web"`
	Token ConfToken `mapstructure:"token"`
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
	Host       string `mapstructure:"host"`  // TODO (PavelS) Remove or redone
}

type ConfWeb struct {
	Server ConfServer `mapstructure:"server"`
	Static ConfStatic `mapstructure:"static"`
}

type ConfToken struct {
	SecretName string `mapstructure:"secretname"`
	CookieName string `mapstructure:"cookiename"`
	Value      int    `mapstructure:"value"`
}

type ConfServer struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type ConfStatic struct {
	DirImg        string `mapstructure:"dir_img"`
	UrlImg        string `mapstructure:"url_img"`
	DirAvt        string `mapstructure:"dir_avt"`
}

func newConfig() *Config {
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
	viper.SetDefault("web.static.dir_img", "img")
	viper.SetDefault("web.static.url_img", "img")
	viper.SetDefault("web.static.dir_avt", "/static/avatar/")
}
