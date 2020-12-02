package postgres

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	repository2 "github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/repository"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
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
var br *repository2.Repository

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
	br = repository2.NewRepo(db)

	code := m.Run()
	os.Exit(code)
}

func TestRepository_Pin(t *testing.T) {
	u, err := ur.CreateUser(&api.SignUp{
		Username: "21savage",
		Password: "123321123",
		Email:    "21@email.com",
	})
	assert.NoError(t, err)

	b, err = br.CreateBoard(u.ID, &api.CreateBoard{
		Title:   "doska",
		Content: "novaya",
	})
	assert.NoError(t, err)

	u.Username = "21savage"

	pin = &domain.Pin{
		Title:   "album",
		Content: "the savage mode",
		UserId:  u.ID,
	}

	r := NewRepo(db)
	err = r.StorePin(pin)
	assert.NoError(t, err)

	p, err := r.GetPin(pin.Id)
	assert.NoError(t, err)
	assert.Equal(t, pin.Title, p.Title)

	_, err = r.GetPin(-1)
	assert.Error(t, err)

	pins, err := r.GetPinList(u.Username)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pins))

	err = br.AttachPin(b.ID, p.Id)
	assert.NoError(t, err)

	pins, err = r.GetPinBoardList(b.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pins))

	_, err = db.Exec(context.Background(), "delete from users where username = '21savage'")
	assert.NoError(t, err)
	_, err = db.Exec(context.Background(), "delete from boards where id = $1", b.ID)
	assert.NoError(t, err)
}
