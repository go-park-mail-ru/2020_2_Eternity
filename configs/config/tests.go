package config

import "github.com/spf13/viper"

func NewConfigTst() *Config {
	setDefaultDbTst()
	setDefaultWebTst()
	setDefaultLogTst()

	conf := new(Config)

	er := viper.Unmarshal(conf)
	if er != nil {
		Lg("config", "newConfig").Fatal("Fatal error config file:", er)
	}

	return conf
}

func setDefaultDbTst() {
	viper.SetDefault("database.postgres.driver_name", "postgres")
	viper.SetDefault("database.postgres.username", "pinterest_user")
	viper.SetDefault("database.postgres.password", "662f2710-4e08-4be7-a278-a53ae86ba7f6")
	viper.SetDefault("database.postgres.db_name", "pinterest_db")
	viper.SetDefault("database.postgres.ssl_mode", "disable")
	viper.SetDefault("database.postgres.host", "localhost")
}

func setDefaultWebTst() {
	viper.SetDefault("web.server.address", "")
	viper.SetDefault("web.server.port", "0000")
	viper.SetDefault("web.server.host", "pinterest-tp.tk")
	viper.SetDefault("web.server.protocol", "http://")
	viper.SetDefault("web.static.dir_img", "img")
	viper.SetDefault("web.static.url_img", "img")
	viper.SetDefault("web.static.dir_avt", "/static/avatar/")
}

func setDefaultLogTst() {
	viper.SetDefault("logger.gin_file", "/var/log/pinterest/gin.log")
	viper.SetDefault("logger.common_file", "/var/log/pinterest/common.log")
	viper.SetDefault("logger.gin_level", "debug")
	viper.SetDefault("logger.common_level", "debug")
	viper.SetDefault("logger.stdout_log", true)
}
