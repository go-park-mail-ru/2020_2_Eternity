package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	Conf = newConfig()
	Db   = newDatabase(&Conf.Db).Open()
)

type Config struct {
	Db     ConfDB     `mapstructure:"database"`
	Web    ConfWeb    `mapstructure:"web"`
	Token  ConfToken  `mapstructure:"token"`
	Logger ConfLogger `mapstructure:"logger"`
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
	Host       string `mapstructure:"host"` // TODO (PavelS) Remove or redone
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
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Protocol string `mapstructure:"protocol"`
}

type ConfStatic struct {
	DirImg string `mapstructure:"dir_img"`
	UrlImg string `mapstructure:"url_img"`
	DirAvt string `mapstructure:"dir_avt"`
}

type ConfLogger struct {
	GinFilePath    string `mapstructure:"gin_file"`
	CommonFilePath string `mapstructure:"common_file"`
	GinLevel       string `mapstructure:"gin_level"`
	CommonLevel    string `mapstructure:"common_level"`
	StdoutLog      bool   `mapstructure:"stdout_log"`
}

func newConfig() *Config {
	setDefaultDb()
	setDefaultWeb()
	setDefaultLog()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	rootDir, exists := os.LookupEnv("ROOT_DIR")
	if exists {
		viper.AddConfigPath(filepath.Join(rootDir, "/configs/yaml"))
	} else {
		viper.AddConfigPath("./configs/yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		Lg("config", "newConfig").Fatal("Fatal error config file ", err)
	}

	conf := new(Config)

	er := viper.Unmarshal(conf)
	if er != nil {
		Lg("config", "newConfig").Fatal("Fatal error config file:", err)
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
	viper.SetDefault("web.server.host", "pinterest-tp.tk")
	viper.SetDefault("web.server.protocol", "http://")
	viper.SetDefault("web.static.dir_img", "img")
	viper.SetDefault("web.static.url_img", "img")
	viper.SetDefault("web.static.dir_avt", "/static/avatar/")
}

func setDefaultLog() {
	viper.SetDefault("logger.gin_file", "/var/log/pinterest/gin.log")
	viper.SetDefault("logger.common_file", "/var/log/pinterest/common.log")
	viper.SetDefault("logger.gin_level", "debug")
	viper.SetDefault("logger.common_level", "debug")
	viper.SetDefault("logger.stdout_log", true)
}
