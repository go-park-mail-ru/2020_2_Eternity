package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	config.Conf = config.NewConfig()
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
	code := m.Run()

	db.Close()
	os.Exit(code)
}


//func TestStore(t *testing.T) {
//	r := NewRepo(db)
//
//	p := domainChat.Chat{}
//
//	assert.Nil(t, r.StoreChat(&p, "name", "name2"))
//	fmt.Println(p)
//
//}



//func TestGetUserChats(t *testing.T) {
//	r := NewRepo(db)
//
//	chats, err := r.GetUserChats("na3")
//
//	assert.Nil(t, err)
//	fmt.Println(chats)
//
//}

//func TestChatById(t *testing.T) {
//	r := NewRepo(db)
//
//	chat, err := r.GetChatById(1, "na")
//
//	assert.Nil(t, err)
//	fmt.Println(chat)
//}

func TestChatMkRead(t *testing.T) {
	r := NewRepo(db)

	err := r.MarkAllMessagesRead(1, "name3")

	assert.Nil(t, err)
}