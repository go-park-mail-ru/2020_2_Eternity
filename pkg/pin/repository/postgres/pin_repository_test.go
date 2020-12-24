package postgres

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

var pin = &domain.Pin{
	Title:         "album drop",
	Content:       "the savage mode",
	Id:            1,
	PictureHeight: 200,
	PictureWidth:  200,
	UserId:        1,
	Username:      "21savage",
}

func TestRepository_CreatePin(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(2)

	mock.ExpectQuery("insert into pins").
		WithArgs(pin.Title, pin.Content, pin.UserId, pin.PictureName, pin.PictureHeight, pin.PictureHeight).
		WillReturnRows(rows)
	err = r.StorePin(pin)
	assert.Equal(t, 2, pin.Id)
	assert.NoError(t, err)

	mock.ExpectQuery("insert into pins").
		WithArgs(pin.Title, pin.Content, pin.UserId, pin.PictureName, pin.PictureHeight, pin.PictureHeight).WillReturnError(errors.New("err"))

	err = r.StorePin(pin)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetPin(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"p.pin_id", "title", "content", "p.name", "user_id", "height", "width", "username", "avatar"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureWidth, pin.Username, pin.Avatar)
	mock.ExpectQuery("select ").WithArgs(1).WillReturnRows(rows)

	p, err := r.GetPin(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, p.Id)

	mock.ExpectQuery("select ").WithArgs(3).WillReturnError(errors.New("not found pin"))
	_, err = r.GetPin(3)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetPinList(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "name", "user_id", "height", "width"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureWidth)
	mock.ExpectQuery("select ").WithArgs(pin.Username).WillReturnRows(rows)

	pins, err := r.GetPinList(pin.Username)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)

	mock.ExpectQuery("select").WithArgs(pin.Username).WillReturnError(errors.New("error"))
	_, err = r.GetPinList(pin.Username)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetBoardPinList(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "name", "user_id", "height", "width"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureWidth)
	mock.ExpectQuery("select ").WithArgs(1).WillReturnRows(rows)

	pins, err := r.GetPinBoardList(1)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)

	mock.ExpectQuery("select").WithArgs(2).WillReturnError(errors.New("error"))
	_, err = r.GetPinBoardList(2)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
