package postgres

//import (
//	"context"
//	"fmt"
//	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/jackc/pgx/v4/pgxpool"
//	"os"
//	"testing"
//)

//var db *pgxpool.Pool
//
//func TestMain(m *testing.M) {
//	config.Conf = config.NewConfig()
//	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
//		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%s",
//		config.Conf.Db.Postgres.Username,
//		config.Conf.Db.Postgres.Password,
//		config.Conf.Db.Postgres.Host,
//		config.Conf.Db.Postgres.DbName,
//		config.Conf.Db.Postgres.SslMode,
//		config.Conf.Db.Postgres.MaxConn,
//	))
//	if err != nil {
//		fmt.Println("Error ", err.Error())
//	}
//
//	db, err = pgxpool.ConnectConfig(context.Background(), conf)
//
//	code := m.Run()
//
//	db.Close()
//	os.Exit(code)
//}


//func TestStore(t *testing.T) {
//	r := NewRepo(db)
//
//	c := domain.Comment{
//		Content: "conte",
//		PinId: 4,
//		UserId: 2,
//	}
//
//	assert.Nil(t, r.StoreChildComment(&c, 42))
//	fmt.Println(c)
//
//}

//func TestStoreRoot(t *testing.T) {
//	r := NewRepo(db)
//
//	c := domain.Comment{
//		Content: "content",
//		PinId: 4,
//		UserId: 2,
//	}
//
//	assert.Nil(t, r.StoreRootComment(&c))
//	fmt.Println(c)
//}


//func TestGetId(t *testing.T) {
//	r := NewRepo(db)
//
//
//	c, err := r.GetComment(53)
//	assert.Nil(t, err)
//	fmt.Println(c)
//}

//func TestPinComments(t *testing.T) {
//	r := NewRepo(db)
//
//
//	c, err := r.GetPinComments(311)
//	assert.Nil(t, err)
//	fmt.Println(c)
//}

