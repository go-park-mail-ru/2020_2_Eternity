package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var user = &domain.User{
	Username: "21savage",
	Avatar:   "asdasda",
	ID:       1,
}

var pin = &domain.Pin{
	Id:            1,
	Title:         "test",
	Content:       "another test",
	PictureName:   "asdasdas",
	PictureHeight: 220,
	PictureWidth:  220,
	UserId:        1,
	Username:      "21savage",
}

var board = &domain.Board{
	ID:       1,
	Title:    "doska",
	Content:  "novaya",
	Username: "21savage",
}

func TestRepository_GetUsersByName(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "username", "avatar"}).AddRow(user.ID, user.Username, user.Avatar)

	mock.ExpectQuery("select ").WithArgs("21sav", 10000).WillReturnRows(rows)
	users, err := r.GetUsersByName("21sav", 10000)
	assert.NoError(t, err)
	assert.Equal(t, len(users), 1)

	mock.ExpectQuery("select ").WithArgs("22sav", 10000).WillReturnError(errors.New("no results"))
	_, err = r.GetUsersByName("22sav", 10000)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetPinsByTitle(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "name", "user_id", "height", "width"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureHeight)

	mock.ExpectQuery("select ").WithArgs("test", 10000).WillReturnRows(rows)
	pins, err := r.GetPinsByTitle("test", 10000)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)

	mock.ExpectQuery("select ").WithArgs("what", 10000).WillReturnError(errors.New("no results"))
	_, err = r.GetPinsByTitle("what", 10000)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetBoardsByTitle(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "username"}).
		AddRow(board.ID, board.Title, board.Content, board.Username)

	mock.ExpectQuery("select ").WithArgs("doska", 10000).WillReturnRows(rows)
	boards, err := r.GetBoardsByTitle("doska", 10000)
	assert.NoError(t, err)
	assert.Equal(t, len(boards), 1)

	mock.ExpectQuery("select ").WithArgs("newd", 10000).WillReturnError(errors.New("no results"))
	_, err = r.GetBoardsByTitle("newd", 10000)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
