package postgres

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	note = domain.Notification{
		Id: 1,
		ToUserId: 2,
		Type: 3,
		EncodedData: []byte("123123"),
		CreationTime: time.Now(),
		IsRead: false,
	}
)

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	code := m.Run()
	os.Exit(code)
}


func TestStoreNote(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectQuery("INSERT INTO").
		WithArgs(n.ToUserId, n.Type, n.EncodedData, sqlmock.AnyArg(), n.IsRead).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time", "is_read"}).
				AddRow(n.Id, n.CreationTime, n.IsRead))


	r := NewRepo(db)
	err := r.StoreNote(&n)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err

	mock.ExpectQuery("INSERT INTO").
		WithArgs(n.ToUserId, n.Type, n.EncodedData, sqlmock.AnyArg(), n.IsRead).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time", "is_read"}))


	er := r.StoreNote(&n)

	assert.NotNil(t, er)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetNoteById(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectQuery("SELECT id").
		WithArgs(n.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "to_user_id", "type", "encoded_data", "creation_time", "is_read"}).
				AddRow(n.Id, n.ToUserId, n.Type, n.EncodedData, n.CreationTime, n.IsRead))


	r := NewRepo(db)
	_, err := r.GetNoteById(n.Id)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err

	mock.ExpectQuery("SELECT id").
		WithArgs(n.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "to_user_id", "type", "encoded_data", "creation_time", "is_read"}))


	_, er := r.GetNoteById(n.Id)

	assert.NotNil(t, er)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetNotes(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectQuery("SELECT id").
		WithArgs(n.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "to_user_id", "type", "encoded_data", "creation_time", "is_read"}).
				AddRow(n.Id, n.ToUserId, n.Type, n.EncodedData, n.CreationTime, n.IsRead))


	r := NewRepo(db)
	_, err := r.GetNoteById(n.Id)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err

	mock.ExpectQuery("SELECT id").
		WithArgs(n.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "to_user_id", "type", "encoded_data", "creation_time", "is_read"}))


	_, er := r.GetNoteById(n.Id)

	assert.NotNil(t, er)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdNoteId(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectExec("UPDATE notifications").
		WithArgs(n.Id).WillReturnResult(sqlmock.NewResult(1, 1))


	r := NewRepo(db)
	err := r.UpdateNoteIsRead(n.Id)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err

	mock.ExpectExec("UPDATE notifications").
		WithArgs(n.Id).
		WillReturnError(errors.New(""))

	er := r.UpdateNoteIsRead(n.Id)

	assert.NotNil(t, er)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdNotes(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectExec("UPDATE notifications").
		WithArgs(n.Id).WillReturnResult(sqlmock.NewResult(1, 1))


	r := NewRepo(db)
	err := r.UpdateUserNotes(n.Id)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err

	mock.ExpectExec("UPDATE notifications").
		WithArgs(n.Id).
		WillReturnError(errors.New(""))

	er := r.UpdateUserNotes(n.Id)

	assert.NotNil(t, er)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetNotesToUser(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	n := note
	mock.ExpectQuery("SELECT id").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "to_user_id", "type", "encoded_data", "creation_time", "is_read"}).
			AddRow(n.Id, n.ToUserId, n.Type, n.EncodedData, n.CreationTime, n.IsRead))

	r := NewRepo(db)
	_, err := r.GetNotesToUser(1)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err


	mock.ExpectQuery("SELECT id").
		WithArgs(1).
		WillReturnError(fmt.Errorf(""))


	_, err = r.GetNotesToUser(1)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}


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

