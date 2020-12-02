package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
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
	code := m.Run()
	os.Exit(code)
}

var u *api.SignUp = &api.SignUp{
	Username: "21savage",
	Email:    "21sav@mail.ru",
	Password: "12345678",
}

func TestRepository_CreateUser(t *testing.T) {
	t.Helper()
	r := NewRepo(db)

	_, err := r.CreateUser(u)
	assert.Equal(t, err, nil)

	_, err = r.CreateUser(u)
	assert.Error(t, err)

	_, err = db.Exec(context.Background(), "delete from users where username = $1", u.Username)
	assert.Equal(t, err, nil)
}

func TestRepository_GetUser(t *testing.T) {
	t.Helper()
	r := NewRepo(db)

	ru, err := r.CreateUser(u)
	assert.Equal(t, err, nil)

	user, err := r.GetUser(ru.ID)
	assert.NoError(t, err)
	assert.Equal(t, u.Username, user.Username)

	user, err = r.GetUserByName("21savage")
	assert.NoError(t, err)
	assert.Equal(t, u.Username, user.Username)

	_, err = r.GetUser(ru.ID - 1)
	assert.Error(t, err)

	user, err = r.GetUserByName("test")
	assert.Error(t, err)

	_, err = db.Exec(context.Background(), "delete from users where username = $1", u.Username)
	assert.Equal(t, err, nil)
}

func TestRepository_UpdateUser(t *testing.T) {
	t.Helper()
	r := NewRepo(db)

	ru, err := r.CreateUser(u)
	assert.Equal(t, err, nil)

	up := &api.UpdateUser{
		Username:    "22savage",
		Email:       "22sav@gmail.com",
		Name:        "Gangster",
		Surname:     "Gangster",
		Description: "4l top album",
	}
	_, err = r.UpdateUser(ru.ID, up)
	assert.NoError(t, err)

	err = r.UpdatePassword(ru.ID, "123456789")
	assert.NoError(t, err)

	err = r.UpdateAvatar(ru.ID, "default.jpg")
	assert.NoError(t, err)

	err, avatar := r.GetAvatar(ru.ID)
	assert.NoError(t, err)
	assert.Equal(t, avatar, "default.jpg")

	err, _ = r.GetAvatar(ru.ID - 15)
	assert.Error(t, err)

	_, err = db.Exec(context.Background(), "delete from users where username = $1", up.Username)
	assert.Equal(t, err, nil)
}

func TestRepository_Follow(t *testing.T) {
	r := NewRepo(db)

	fu, err := r.CreateUser(u)
	assert.Equal(t, err, nil)

	su, err := r.CreateUser(&api.SignUp{
		Username: "future",
		Password: "12345678",
		Email:    "future@mail.ru",
	})
	assert.Equal(t, err, nil)

	err = r.Follow(fu.ID, su.ID)
	assert.NoError(t, err)
	err = r.Follow(fu.ID, su.ID)
	assert.Error(t, err)

	users, err := r.GetFollowers("future")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))

	arr, err := r.GetFollowersIds(su.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(arr))

	second, err := r.GetUserByNameWithFollowers("future")
	assert.NoError(t, err)
	assert.Equal(t, second.Followers, 1)

	_, err = r.GetUserByNameWithFollowers("test")
	assert.Error(t, err)

	users, err = r.GetFollowing(u.Username)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))

	err = r.IsFollowing(fu.ID, second.Username)
	assert.NoError(t, err)

	err = r.IsFollowing(su.ID, u.Username)
	assert.Error(t, err)

	err = r.UnFollow(fu.ID, su.ID)
	assert.NoError(t, err)

	_, err = db.Exec(context.Background(), "delete from users where username = $1", u.Username)
	assert.Equal(t, err, nil)
	_, err = db.Exec(context.Background(), "delete from users where username = $1", second.Username)
	assert.Equal(t, err, nil)
}
