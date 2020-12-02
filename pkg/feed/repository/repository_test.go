package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	pin_postgres "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/repository/postgres"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var db *pgxpool.Pool

var u *domain.User
var pin *domain.Pin
var b *domain.Board
var ur *repository.Repository
var pr *pin_postgres.Repository

func TestMain(m *testing.M) {
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%d",
		config.Conf.Db.Postgres.Username,
		config.Conf.Db.Postgres.Password,
		config.Conf.Db.Postgres.Host,
		config.Conf.Db.Postgres.DbName,
		config.Conf.Db.Postgres.SslMode,
		10,
	))
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	db, err = pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	ur = repository.NewRepo(db)
	pr = pin_postgres.NewRepo(db)

	code := m.Run()
	os.Exit(code)
}

func TestRepository_Feed(t *testing.T) {
	u, err := ur.CreateUser(&api.SignUp{
		Username: "21savage",
		Password: "123321123",
		Email:    "21@email.com",
	})
	assert.NoError(t, err)

	u.Username = "21savage"

	pin = &domain.Pin{
		Title:   "album drop",
		Content: "the savage mode",
		UserId:  u.ID,
	}

	r := NewRepo(db)
	err = pr.StorePin(pin)
	assert.NoError(t, err)

	pins, err := r.GetFeed(0, 10000000)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pins))

	_, err = db.Exec(context.Background(), "delete from users where username = '21savage'")
	assert.NoError(t, err)
}
