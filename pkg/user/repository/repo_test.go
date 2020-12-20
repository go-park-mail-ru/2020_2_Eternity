package repository

import (
	"errors"
	//"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var u = &api.SignUp{
	Username: "21savage",
	Email:    "21sav@mail.ru",
	Password: "12345678",
}

func TestRepository_CreateUser(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery("insert into users ").WithArgs(u.Username, u.Email, u.Password,
		u.Name, u.Surname, u.Description, u.BirthDate, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	r := NewRepo(db)
	ur, err := r.CreateUser(u)
	assert.NoError(t, err)
	assert.Equal(t, 1, ur.ID)

	mock.ExpectQuery("insert into users ").WithArgs(u.Username, u.Email, u.Password,
		u.Name, u.Surname, u.Description, u.BirthDate, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("user exists"))

	_, err = r.CreateUser(u)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetUser(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

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
}

/*func TestRepository_UpdateUser(t *testing.T) {
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
*/
