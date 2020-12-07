package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*var db *pgxpool.Pool

func TestMain(m *testing.M) {
	config.Conf = config.NewConfig()
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%s",
		config.Conf.Db.Postgres.Username,
		config.Conf.Db.Postgres.Password,
		config.Conf.Db.Postgres.Host,
		config.Conf.Db.Postgres.DbName,
		config.Conf.Db.Postgres.SslMode,
		config.Conf.Db.Postgres.MaxConn,
	))
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	db, err = pgxpool.ConnectConfig(context.Background(), conf)
	code := m.Run()

	db.Close()
	os.Exit(code)
}*/

//func TestStore(t *testing.T) {
//	r := NewRepo(db)
//
//	p := domainChat.Chat{}
//
//	assert.Nil(t, r.StoreChat(&p, 2, "name2"))
//	fmt.Println(p)
//
//}

//func TestGetUserChats(t *testing.T) {
//	r := NewRepo(db)
//
//	chats, err := r.GetUserChats(2)
//
//	assert.Nil(t, err)
//	fmt.Println(chats)
//
//}

//func TestChatById(t *testing.T) {
//	r := NewRepo(db)
//
//	chat, err := r.GetChatById(1, 2)
//
//	assert.Nil(t, err)
//	fmt.Println(chat)
//}

//func TestChatMkRead(t *testing.T) {
//	r := NewRepo(db)
//
//	err := r.MarkAllMessagesRead(1, 10)
//
//	assert.Nil(t, err)
//}

//func TestStoreMsg(t *testing.T) {
//	r := NewRepo(db)
//
//	mReq := domainChat.CreateMessageReq{
//		ChatId: 1,
//		Content: "jopa",
//	}
//
//	m, err := r.StoreMessage(&mReq, 2)
//
//	fmt.Println(m)
//
//	assert.Nil(t, err)
//}

//func TestDeleteMsg(t *testing.T) {
//	r := NewRepo(db)
//
//
//	err := r.DeleteMessage(3)
//
//	assert.Nil(t, err)
//}

//func TestGetLastNMsgs(t *testing.T) {
//	r := NewRepo(db)
//
//	mReq := domainChat.GetLastNMessagesReq{
//		ChatId: 1,
//		NMessages: 6,
//	}
//
//	m, err := r.GetLastNMessages(&mReq)
//
//	fmt.Println(m)
//
//	assert.Nil(t, err)
//}

//func TestGetLastNMsgs(t *testing.T) {
//	r := NewRepo(db)
//
//	mReq := domainChat.GetNMessagesBeforeReq{
//		ChatId: 1,
//		NMessages: 2,
//		BeforeMessageId: 8,
//	}
//
//	m, err := r.GetNMessagesBefore(&mReq)
//
//	fmt.Println(m)
//
//	assert.Nil(t, err)
//}

func TestGeCollId(t *testing.T) {
	r := NewRepo(db)

	id, err := r.GetCollocutorId(2, 1)

	fmt.Println(id)

	assert.Nil(t, err)
}
