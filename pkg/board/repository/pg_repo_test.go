package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
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

var b *domain.Board

var pin = &domain.Pin{
	Title:    "album",
	Content:  "the savage mode",
	Username: "21savage",
	UserId:   1,
}

var cboard = &api.CreateBoard{
	Title:   "doska",
	Content: "novaya",
}

var board = &domain.Board{
	ID:       1,
	Title:    "doska",
	Content:  "novaya",
	Username: "21savage",
}

func TestRepository_Board(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	rowsU := sqlmock.NewRows([]string{"username"}).AddRow("21savage")
	mock.ExpectQuery("insert into boards").WithArgs(cboard.Title, cboard.Content, 1).WillReturnRows(rows)
	mock.ExpectQuery("select ").WithArgs(1).WillReturnRows(rowsU)

	r := NewRepo(db)
	b, err = r.CreateBoard(1, cboard)
	assert.NoError(t, err)
	assert.Equal(t, b.Title, "doska")

	mock.ExpectQuery("insert into boards").WithArgs(cboard.Title, cboard.Content, 2).
		WillReturnError(errors.New("board err"))
	b, err = r.CreateBoard(2, cboard)
	assert.Error(t, err)

	rows2 := sqlmock.NewRows([]string{"id"}).AddRow(2)
	mock.ExpectQuery("insert into boards").WithArgs(cboard.Title, cboard.Content, 4).WillReturnRows(rows2)
	mock.ExpectQuery("select ").WithArgs(4).WillReturnError(errors.New("user not found"))
	b, err = r.CreateBoard(4, cboard)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetBoard(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"title", "content", "username"}).AddRow(board.Title, board.Content, board.Username)
	mock.ExpectQuery("select ").WithArgs(1).WillReturnRows(rows)

	b, err := r.GetBoard(1)
	assert.NoError(t, err)
	assert.Equal(t, b.Content, board.Content)

	mock.ExpectQuery("select ").WithArgs(2).WillReturnError(errors.New("not found board"))
	_, err = r.GetBoard(2)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetAllBoardsByUser(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"board.id", "title", "content"}).AddRow(board.ID, board.Title, board.Content)
	mock.ExpectQuery("select ").WithArgs(board.Username).WillReturnRows(rows)

	boards, err := r.GetAllBoardsByUser(board.Username)

	assert.NoError(t, err)
	assert.Equal(t, len(boards), 1)

	mock.ExpectQuery("select ").WithArgs("22sav").WillReturnError(errors.New("user not found"))

	_, err = r.GetAllBoardsByUser("22sav")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_CheckOwner(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock.ExpectQuery("select ").WithArgs(1).WillReturnRows(rows)
	err = r.CheckOwner(1, board.ID)
	assert.NoError(t, err)

	rows2 := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock.ExpectQuery("select ").WithArgs(2).WillReturnRows(rows2)

	err = r.CheckOwner(2, 2)
	assert.Error(t, err)

	mock.ExpectQuery("select ").WithArgs(3).WillReturnError(errors.New("board not found"))

	err = r.CheckOwner(1, 3)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_AttachPin(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRepo(db)

	mock.ExpectExec("insert into boards_pins").WithArgs(board.ID, pin.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err = r.AttachPin(board.ID, pin.Id)
	assert.NoError(t, err)

	mock.ExpectExec("insert into boards_pins").WithArgs(board.ID, pin.Id).WillReturnError(errors.New("already attach"))
	err = r.AttachPin(board.ID, pin.Id)
	assert.Error(t, err)

	mock.ExpectExec("delete from boards_pins").WithArgs(board.ID, pin.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err = r.DetachPin(board.ID, pin.Id)
	assert.NoError(t, err)

	mock.ExpectExec("delete from boards_pins").WithArgs(board.ID, pin.Id).WillReturnError(errors.New("already detach"))
	err = r.DetachPin(board.ID, pin.Id)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
