package repository

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"time"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}

// Chats

func (r *Repository) StoreChat(ch *domainChat.Chat, userId int, collocutorName string) error {
	tx, err := r.dbConn.Begin()
	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	defer func() {
		err := tx.Rollback()
		if err != nil {
			config.Lg("chat_repo", "StoreChat").Error(err.Error())
		}
	}()

	err = tx.QueryRow(
		"INSERT INTO chats (creation_time) "+
			"VALUES ($1) "+
			"RETURNING id, creation_time; ",
		time.Now()).Scan(&ch.Id, &ch.CreationTime)

	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	_, err = tx.Exec(
		"WITH col_id AS (SELECT id FROM users WHERE username = $1) "+
			"INSERT INTO uu_chat (user_id, collocutor_id, chat_id,  last_read_msg_id, new_messages) "+
			"VALUES "+
			"($2, (SELECT id FROM col_id), $3 , 0, 0), "+
			"((SELECT id FROM col_id), $2, $3 , 0, 0); ",
		collocutorName, userId, ch.Id)

	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	if err = tx.QueryRow(
		"SELECT username FROM users WHERE id = $1",
		userId).Scan(&ch.UserName); err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	ch.CollocutorName = collocutorName

	return nil
}

func (r *Repository) GetChatById(chatId int, userId int) (domainChat.Chat, error) {
	c := domainChat.Chat{}
	err := r.dbConn.QueryRow(
		"WITH usr AS (SELECT username FROM users WHERE id = $1)"+
			"SELECT c.id, c.creation_time, c.last_msg_id, c.last_msg_content, c.last_msg_username, c.last_msg_time, "+
			"(SELECT username FROM usr), "+
			" u.username, u.avatar, uc.last_read_msg_id, uc.new_messages "+
			"FROM "+
			"chats c JOIN uu_chat uc "+
			"ON c.id = uc.chat_id "+
			"JOIN users u "+
			"ON u.id = uc.collocutor_id "+
			"WHERE uc.user_id = $1 AND c.id = $2 ",
		userId, chatId).
		Scan(&c.Id, &c.CreationTime, &c.LastMsgId, &c.LastMsgContent, &c.LastMsgUsername,
			&c.LastMsgTime, &c.UserName, &c.CollocutorName, &c.CollocutorAvatarLink, &c.LastReadMsgId, &c.NewMessages)

	if err != nil {
		config.Lg("chat_repo", "GetChatById").Error(err.Error())
		return domainChat.Chat{}, err
	}

	return c, nil
}

func (r *Repository) GetUserChats(userId int) ([]domainChat.Chat, error) {
	tx, err := r.dbConn.Begin()
	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	defer func() {
		err := tx.Rollback()
		if err != nil {
			config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		}
	}()

	records := 0
	err = tx.QueryRow(
		"SELECT COUNT(1) FROM users WHERE id = $1", userId).Scan(&records)

	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	if records == 0 {
		config.Lg("chat_repo", "GetUserChats").Error(errors.New("User doesn't exist"))
		return nil, errors.New("User doesn't exist")
	}

	chats, err := r.getUserChatsInternal(tx, userId)
	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	return chats, nil
}

func (r *Repository) getUserChatsInternal(tx *sql.Tx, userId int) ([]domainChat.Chat, error) {
	rows, err := (*tx).Query(
		"SELECT c.id, c.creation_time, c.last_msg_id, c.last_msg_content, c.last_msg_username, "+
			"c.last_msg_time, u_user.username, u_col.username, u_col.avatar, uc.last_read_msg_id, uc.new_messages "+
			"FROM "+
			"chats c JOIN uu_chat uc "+
			"ON c.id = uc.chat_id "+
			"JOIN users u_col "+
			"ON u_col.id = uc.collocutor_id "+
			"JOIN users u_user "+
			"ON u_user.id = uc.user_id "+
			"WHERE uc.user_id = $1 "+
			"ORDER BY c.last_msg_time DESC, c.creation_time DESC ",
		userId)

	if err != nil {
		config.Lg("chat_repo", "GetUserChatsInt").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	chats := []domainChat.Chat{}
	for rows.Next() {
		c := domainChat.Chat{}
		err := rows.Scan(&c.Id, &c.CreationTime, &c.LastMsgId, &c.LastMsgContent, &c.LastMsgUsername,
			&c.LastMsgTime, &c.UserName, &c.CollocutorName, &c.CollocutorAvatarLink, &c.LastReadMsgId, &c.NewMessages)

		if err != nil {
			config.Lg("comment_postgres", "GetUserChatsInt").Error(err.Error())
			return nil, err
		}

		chats = append(chats, c)
	}

	if rows.Err() != nil {
		config.Lg("comment_postgres", "GetUserChatsInt").Error(rows.Err())
		return nil, rows.Err()
	}

	return chats, nil
}

func (r *Repository) MarkAllMessagesRead(chatId int, userId int) error {
	_, err := r.dbConn.Exec(
		"UPDATE uu_chat "+
			"SET new_messages = 0, "+
			"last_read_msg_id = chats.last_msg_id "+
			"FROM chats "+
			"WHERE chat_id = chats.id AND "+
			"user_id = $1 AND "+
			"chat_id = $2 ",
		userId, chatId)

	if err != nil {
		config.Lg("chat_repo", "GetChatById").Error(err.Error())
		return err
	}

	return nil
}

// Messages

func (r *Repository) StoreMessage(mReq *domainChat.CreateMessageReq, userId int) (domainChat.Message, error) {
	// TODO (Pavel S) Add protection of writing to foreign chat
	m := domainChat.Message{}
	err := r.dbConn.QueryRow(
		"WITH u AS (SELECT username, avatar FROM users WHERE id = $1) "+
			"INSERT INTO messages (content, creation_time, chat_id, user_id, username, avatar) "+
			"VALUES ($2, $3, $4, $1, (SELECT username FROM u), (SELECT avatar FROM u)) "+
			"RETURNING id, content, creation_time, chat_id, user_id, username, avatar ",
		userId, mReq.Content, time.Now(), mReq.ChatId).
		Scan(&m.Id, &m.Content, &m.CreationTime, &m.ChatId, &m.UserId, &m.UserName, &m.UserAvatarLink)

	if err != nil {
		config.Lg("chat_repo", "StoreMessage").Error(err.Error())
		return domainChat.Message{}, err
	}

	return m, nil
}

func (r *Repository) DeleteMessage(msgId int) error {
	_, err := r.dbConn.Exec("DELETE FROM messages WHERE id = $1 ", msgId)

	if err != nil {
		config.Lg("chat_repo", "DeleteMessage").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.Message, error) {
	rows, err := r.dbConn.Query(
		"SELECT * FROM "+
			"(SELECT id, content, creation_time, chat_id, user_id, username, avatar "+
			"FROM messages "+
			"WHERE chat_id = $1 "+
			"ORDER BY id DESC "+
			"LIMIT $2) t "+
			"ORDER BY id DESC",
		mReq.ChatId, mReq.NMessages)

	if err != nil {
		config.Lg("chat_repo", "GetLastNMessages").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	msgs := []domainChat.Message{}
	for rows.Next() {
		m := domainChat.Message{}
		err := rows.Scan(&m.Id, &m.Content, &m.CreationTime, &m.ChatId, &m.UserId, &m.UserName, &m.UserAvatarLink)

		if err != nil {
			config.Lg("comment_postgres", "GetLastNMessages").Error(err.Error())
			return nil, err
		}

		msgs = append(msgs, m)
	}

	if rows.Err() != nil {
		config.Lg("comment_postgres", "GetLastNMessages").Error(rows.Err())
		return nil, rows.Err()
	}

	return msgs, nil
}

func (r *Repository) GetNMessagesBefore(mReq *domainChat.GetNMessagesBeforeReq) ([]domainChat.Message, error) {
	rows, err := r.dbConn.Query(
		"SELECT id, content, creation_time, chat_id, user_id, username, avatar "+
			"FROM messages "+
			"WHERE chat_id = $1 AND id < $2 "+
			"ORDER BY id DESC "+
			"LIMIT $3 ",
		mReq.ChatId, mReq.BeforeMessageId, mReq.NMessages)

	if err != nil {
		config.Lg("chat_repo", "GetNMessagesBefore").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	msgs := []domainChat.Message{}
	for rows.Next() {
		m := domainChat.Message{}
		err := rows.Scan(&m.Id, &m.Content, &m.CreationTime, &m.ChatId, &m.UserId, &m.UserName, &m.UserAvatarLink)

		if err != nil {
			config.Lg("comment_postgres", "GetNMessagesBefore").Error(err.Error())
			return nil, err
		}

		msgs = append(msgs, m)
	}

	if rows.Err() != nil {
		config.Lg("comment_postgres", "GetNMessagesBefore").Error(rows.Err())
		return nil, rows.Err()
	}

	return msgs, nil
}

func (r *Repository) GetCollocutorId(userId int, chatId int) (int, error) {
	collocutorId := 0
	if err := r.dbConn.QueryRow(
		"SELECT collocutor_id FROM uu_chat WHERE user_id = $1 AND chat_id = $2 ",
		userId, chatId).
		Scan(&collocutorId); err != nil {
		config.Lg("chat_repo", "GetCollocutorId").Error(err.Error())
		return 0, err
	}

	return collocutorId, nil
}
