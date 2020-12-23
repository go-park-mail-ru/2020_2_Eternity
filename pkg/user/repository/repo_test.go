package repository

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/lib/pq"
	"time"

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

var user = &domain.User{
	ID:          1,
	Username:    "21savage",
	Email:       "21sav@mail.ru",
	Password:    "12345678",
	Name:        "next",
	Surname:     "gen",
	Description: "rap",
	BirthDate:   time.Time{},
	Avatar:      "123214124",
	Followers:   0,
	Following:   0,
}

var upUs = &api.UpdateUser{
	Username:    "22savage",
	Email:       "22sav@mail.ru",
	Name:        "gen",
	Surname:     "next",
	Description: "12321",
	BirthDate:   time.Time{},
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

	mock.ExpectQuery("insert into users ").WithArgs(user.Username, u.Email, u.Password,
		u.Name, u.Surname, u.Description, u.BirthDate, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	r := NewRepo(db)
	ur, err := r.CreateUser(u)
	assert.NoError(t, err)
	assert.Equal(t, 1, ur.ID)

	mock.ExpectQuery("insert into users ").WithArgs(u.Username, u.Email, u.Password,
		u.Name, u.Surname, u.Description, u.BirthDate, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("user exists"))

	_, _ = r.CreateUser(u)
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

	rows := sqlmock.NewRows([]string{"username", "name", "surname", "description", "password", "email", "birthdate", "avatar"}).
		AddRow(u.Username, u.Name, u.Surname, u.Description, u.Password, u.Email, u.BirthDate, user.Avatar)

	mock.ExpectQuery("select").WithArgs(1).WillReturnRows(rows)

	user, err := r.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, u.Username, user.Username)

	mock.ExpectQuery("select ").WithArgs(1).
		WillReturnError(errors.New("not found"))
	_, err = r.GetUser(1)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetUserByName(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email", "birthdate", "avatar"}).
		AddRow(user.ID, user.Username, user.Password, user.Email, user.BirthDate, user.Avatar)

	mock.ExpectQuery("select").WithArgs(u.Username).WillReturnRows(rows)

	user, err := r.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, u.Username, user.Username)

	mock.ExpectQuery("select ").WithArgs(u.Username).
		WillReturnError(errors.New("not found"))
	_, err = r.GetUserByName(user.Username)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_UpdateUser(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	mock.ExpectExec("update users set").WithArgs(upUs.Email, upUs.Username,
		upUs.Name, upUs.Surname, upUs.Description, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := r.UpdateUser(1, upUs)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, upUs.Username)

	mock.ExpectExec("update users set ").WithArgs(upUs.Email, upUs.Username, upUs.Name,
		upUs.Surname, upUs.Description, 1).
		WillReturnError(errors.New("not found"))
	_, err = r.UpdateUser(1, upUs)
	assert.Error(t, err)

	// update password

	pswds := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123456789",
	}

	mock.ExpectExec("update users set").WithArgs(pswds.NewPassword, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.UpdatePassword(1, pswds.NewPassword)
	assert.NoError(t, err)

	mock.ExpectExec("update users set").WithArgs(pswds.OldPassword, 1).
		WillReturnError(errors.New("user not found"))

	err = r.UpdatePassword(1, pswds.OldPassword)
	assert.Error(t, err)

	// update avatar
	mock.ExpectExec("update users set ").WithArgs("new_avatar", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.UpdateAvatar(1, "new_avatar")
	assert.NoError(t, err)

	mock.ExpectExec("update users set").WithArgs("bad_avatar", 1).
		WillReturnError(errors.New("user not found"))

	err = r.UpdateAvatar(1, "bad_avatar")
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetAvatar(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("avatar")
	mock.ExpectQuery("select avatar from users").WithArgs(1).WillReturnRows(rows)

	err, ava := r.GetAvatar(1)
	assert.NoError(t, err)
	assert.Equal(t, "avatar", ava)

	mock.ExpectQuery("select avatar from users").WithArgs(1).WillReturnError(errors.New("user not found"))
	err, _ = r.GetAvatar(1)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_Follow(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	mock.ExpectExec("insert into follows").WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.Follow(1, 2)
	assert.NoError(t, err)

	mock.ExpectExec("insert into follows").WithArgs(1, 2).
		WillReturnError(errors.New("already follows"))
	err = r.Follow(1, 2)
	assert.Error(t, err)

	//unfollow

	mock.ExpectExec("delete from follows").WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.UnFollow(1, 2)
	assert.NoError(t, err)

	mock.ExpectExec("delete from follows").WithArgs(1, 2).
		WillReturnError(errors.New("already unfollows"))
	err = r.UnFollow(1, 2)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetUserByNameWithFollowers(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "surname", "description", "avatar", "followers", "following"}).
		AddRow(user.ID, user.Username, user.Password, user.Name, user.Surname, user.Description, user.Avatar, user.Followers, user.Following)

	mock.ExpectQuery("select").WithArgs(u.Username).WillReturnRows(rows)

	user, err := r.GetUserByNameWithFollowers(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, u.Username, user.Username)

	mock.ExpectQuery("select ").WithArgs(u.Username).
		WillReturnError(errors.New("not found"))
	_, err = r.GetUserByNameWithFollowers(user.Username)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetFollowersIds(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)
	followersSql := []sql.NullInt32{{Int32: 3, Valid: true}, {Int32: 2, Valid: true}}
	rows := sqlmock.NewRows([]string{"array"}).AddRow(pq.Array(followersSql))

	mock.ExpectQuery("select ARRAY").WithArgs(1).WillReturnRows(rows)
	arr, err := r.GetFollowersIds(1)

	assert.NoError(t, err)
	assert.Equal(t, len(arr), 2)

	mock.ExpectQuery("select ARRAY").WithArgs(1).WillReturnError(errors.New("not found"))
	_, err = r.GetFollowersIds(1)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetFollowers(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"username", "avatar"}).AddRow("asaprocky", "asdqew").AddRow("future", "sadq12")

	mock.ExpectQuery("select").WithArgs(user.Username).WillReturnRows(rows)

	f, err := r.GetFollowers(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, len(f), 2)

	mock.ExpectQuery("select").WithArgs(user.Username).WillReturnError(errors.New("not found"))
	_, err = r.GetFollowers(user.Username)
	assert.Error(t, err)

	rows2 := sqlmock.NewRows([]string{"username", "avatar"}).AddRow("asaprocky", "asdqew").AddRow("future", "sadq12")

	mock.ExpectQuery("select").WithArgs(user.Username).WillReturnRows(rows2)

	f, err = r.GetFollowing(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, len(f), 2)

	mock.ExpectQuery("select").WithArgs(user.Username).WillReturnError(errors.New("not found"))
	_, err = r.GetFollowing(user.Username)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_IsFollowing(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id2"}).AddRow(1)
	mock.ExpectQuery("select ").WithArgs(user.Username, 2).WillReturnRows(rows)
	err = r.IsFollowing(2, user.Username)
	assert.NoError(t, err)

	mock.ExpectQuery("select ").WithArgs(user.Username, 3).WillReturnError(errors.New("not found user"))
	err = r.IsFollowing(3, user.Username)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

/*
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
