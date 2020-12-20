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

var pin = &domain.Pin{
	Title:         "album drop",
	Content:       "the savage mode",
	Id:            1,
	PictureHeight: 200,
	PictureWidth:  200,
	UserId:        1,
	Username:      "21savage",
}

func TestRepository_Feed(t *testing.T) {

	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"pins.id", "title", "content", "name", "user_id", "height", "width"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureWidth)
	mock.ExpectQuery("select ").WithArgs(10000).WillReturnRows(rows)

	pins, err := r.GetFeed(0, 10000)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)

	mock.ExpectQuery("select").WithArgs(10000).WillReturnError(errors.New("feed"))
	_, err = r.GetFeed(0, 10000)
	assert.Error(t, err)

	rows2 := sqlmock.NewRows([]string{"id", "title", "content", "name", "user_id", "height", "width"}).
		AddRow(pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId, pin.PictureHeight, pin.PictureWidth)

	mock.ExpectQuery("select ").WithArgs(1, 10000).WillReturnRows(rows2)

	pins, err = r.GetSubFeed(1, 10000)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)

	mock.ExpectQuery("select").WithArgs(2, 10000).WillReturnError(errors.New("feed"))
	_, err = r.GetSubFeed(2, 10000)
	assert.Error(t, err)
}
