package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	userId = 0
	ch = domainChat.Chat{
		Id: 1,
		CreationTime: time.Now(),
		LastMsgId: 2,
		LastMsgContent: "content",
		LastMsgUsername: "username",
		LastMsgTime: time.Now(),
		UserName: "username",
		CollocutorName: "collocutor",
		CollocutorAvatarLink: "ava",
		LastReadMsgId: 3,
		NewMessages: 5,
	}
	c = domainChat.Chat{
		Id: 1,
		CreationTime: time.Now(),
		LastMsgId: 2,
		LastMsgContent: "content",
		LastMsgUsername: "username",
		LastMsgTime: time.Now(),
		UserName: "username",
		CollocutorName: "collocutor",
		CollocutorAvatarLink: "ava",
		LastReadMsgId: 3,
		NewMessages: 5,
	}
)

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	code := m.Run()
	os.Exit(code)
}


func TestStoreChat(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO chats").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time"}).
				AddRow(ch.Id, ch.CreationTime))
	mock.ExpectExec("WITH col_id AS").
		WithArgs(ch.CollocutorName, userId, ch.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT username").
		WithArgs(userId).
		WillReturnRows(
			sqlmock.NewRows([]string{"username"}).
				AddRow(&ch.UserName))
	mock.ExpectCommit()

	r := NewRepo(db)
	err := r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err begin

	mock.ExpectBegin().WillReturnError(fmt.Errorf(""))

	err = r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err 1 query

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO chats").
		WithArgs(sqlmock.AnyArg()).WillReturnError(fmt.Errorf(""))

	err = r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err 2 query

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO chats").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time"}).
				AddRow(ch.Id, ch.CreationTime))
	mock.ExpectExec("WITH col_id AS").
		WithArgs(ch.CollocutorName, userId, ch.Id).
		WillReturnError(fmt.Errorf(""))

	err = r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err 3 query

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO chats").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time"}).
				AddRow(ch.Id, ch.CreationTime))
	mock.ExpectExec("WITH col_id AS").
		WithArgs(ch.CollocutorName, userId, ch.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT username").
		WithArgs(userId).
		WillReturnError(fmt.Errorf(""))

	err = r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Err commit

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO chats").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "creation_time"}).
				AddRow(ch.Id, ch.CreationTime))
	mock.ExpectExec("WITH col_id AS").
		WithArgs(ch.CollocutorName, userId, ch.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT username").
		WithArgs(userId).
		WillReturnRows(
			sqlmock.NewRows([]string{"username"}).
				AddRow(&ch.UserName))
	mock.ExpectCommit().WillReturnError(fmt.Errorf(""))


	err = r.StoreChat(&ch, userId, ch.CollocutorName)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}


func TestGetById(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	mock.ExpectQuery("WITH usr AS ").
		WithArgs(userId, c.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"c.id", "c.creation_time", "c.last_msg_id",
				"c.last_msg_content", "c.last_msg_username", "c.last_msg_time",
				"(SELECT username FROM usr)",
				" u.username", "u.avatar", "uc.last_read_msg_id", "uc.new_messages"}).
				AddRow(c.Id, c.CreationTime, c.LastMsgId, c.LastMsgContent, c.LastMsgUsername,
				c.LastMsgTime, c.UserName, c.CollocutorName, c.CollocutorAvatarLink, c.LastReadMsgId, c.NewMessages))


	r := NewRepo(db)
	_, err := r.GetChatById(c.Id, userId)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Fail

	mock.ExpectQuery("WITH usr AS ").
		WithArgs(userId, c.Id).
		WillReturnError(fmt.Errorf(""))


	_, err = r.GetChatById(c.Id, userId)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestGetChats(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Ok

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"COUNT(1)"}).AddRow(1))
	mock.ExpectQuery("SELECT c.id").
		WithArgs(userId).
		WillReturnRows(
			sqlmock.NewRows([]string{"c.id", "c.creation_time", "c.last_msg_id",
				"c.last_msg_content", "c.last_msg_username", "c.last_msg_time",
				"(SELECT username FROM usr)",
				" u.username", "u.avatar", "uc.last_read_msg_id", "uc.new_messages"}).
				AddRow(c.Id, c.CreationTime, c.LastMsgId, c.LastMsgContent, c.LastMsgUsername,
					c.LastMsgTime, c.UserName, c.CollocutorName, c.CollocutorAvatarLink, c.LastReadMsgId, c.NewMessages))
	mock.ExpectCommit()

	r := NewRepo(db)
	_, err := r.GetUserChats(userId)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	// Fail

	//mock.ExpectQuery("WITH usr AS ").
	//	WithArgs(userId, c.Id).
	//	WillReturnError(fmt.Errorf(""))
	//
	//
	//_, err = r.GetChatById(c.Id, userId)
	//
	//assert.NotNil(t, err)
	//assert.Nil(t, mock.ExpectationsWereMet())

}