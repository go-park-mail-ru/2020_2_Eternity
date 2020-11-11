package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewTestConfig()
	return true
}()

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%s",
		config.Conf.Db.Postgres.Username,
		config.Conf.Db.Postgres.Password,
		config.Conf.Db.Postgres.Host,
		config.Conf.Db.Postgres.DbName,
		config.Conf.Db.Postgres.SslMode,
		config.Conf.Db.Postgres.MaxConn,
	))
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	db, err = pgxpool.ConnectConfig(context.Background(), conf)
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

func TestRepository_CreateUser(t *testing.T) {

}
