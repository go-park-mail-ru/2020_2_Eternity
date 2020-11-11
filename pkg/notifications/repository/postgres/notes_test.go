package postgres

import (
	"context"
	"encoding/json"
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
//	r := repo.NewRepo(db)
//
//	n := domain.Notification{
//		ToUserId: 3,
//		Type: 3,
//		EncodedData: []byte("jopa"),
//	}
//
//	assert.Nil(t, r.StoreNote(&n))
//	fmt.Println(n)
//	fmt.Println(string(n.EncodedData))
//}

func TestGetId(t *testing.T) {
	r := NewRepo(db)

	n, err := r.GetNoteById(1)
	assert.Nil(t, err)

	fmt.Println(n)
}

func TestGetUserNotes(t *testing.T) {
	r := NewRepo(db)

	ns, err := r.GetNotesToUser(3)
	assert.Nil(t, err)

	fmt.Println(ns)

	type Tst struct {
		B []byte `json:"b"`
	}

	res, _ := json.Marshal(Tst{[]byte("jopa")})
	fmt.Println(string(res))
}

func TestUpdate(t *testing.T) {
	r := NewRepo(db)

	err := r.UpdateNoteIsRead(1)
	assert.Nil(t, err)

}

func TestUpdateAll(t *testing.T) {
	r := NewRepo(db)

	err := r.UpdateUserNotes(3)
	assert.Nil(t, err)

}
