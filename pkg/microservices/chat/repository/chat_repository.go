package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/jackc/pgx/v4"
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

func (r *Repository) StoreChat(ch *domainChat.Chat, userName string, collocutorName string) error {
	tx, err := r.dbConn.Begin(context.Background())
	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(
		context.Background(),
		"INSERT INTO chats (creation_time) " +
			"VALUES ($1) " +
			"RETURNING id, creation_time; ",
			time.Now()).Scan(&ch.Id, &ch.CreationTime)

	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	_, err = tx.Exec(
		context.Background(),
				"WITH cr_id AS (SELECT id FROM users WHERE username = $1), " +
					 "col_id AS (SELECT id FROM users WHERE username = $2) " +
				"INSERT INTO uu_chat (user_id, collocutor_id, chat_id,  last_read_msg_id, new_messages) " +
				"VALUES " +
					"((SELECT id FROM cr_id), (SELECT id FROM col_id), $3 , 0, 0), " +
					"((SELECT id FROM col_id), (SELECT id FROM cr_id), $3 , 0, 0); ",
		userName, collocutorName, ch.Id)


	if err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		config.Lg("chat_repo", "StoreChat").Error(err.Error())
		return err
	}

	ch.CollocutorName = collocutorName
	ch.UserName = userName

	return nil
}



func (r *Repository) GetChatById(chatId int, userName string) (domainChat.Chat, error) {
	c := domainChat.Chat{}
	err := r.dbConn.QueryRow(
		context.Background(),
		"SELECT c.id, c.creation_time, c.last_msg_id, c.last_msg_content, c.last_msg_username, " +
			"c.last_msg_time, u.username, u.avatar, uc.last_read_msg_id, uc.new_messages " +
			"FROM " +
				"chats c JOIN uu_chat uc " +
				"ON c.id = uc.chat_id " +
					"JOIN users u " +
					"ON u.id = uc.collocutor_id " +
			"WHERE uc.user_id = (SELECT id FROM users WHERE username = $1) AND c.id = $2 ",
		userName, chatId).
			Scan(&c.Id, &c.CreationTime, &c.LastMsgId, &c.LastMsgContent, &c.LastMsgUsername,
		&c.LastMsgTime, &c.CollocutorName, &c.CollocutorAvatarLink, &c.LastReadMsgId, &c.NewMessages)

	if err != nil {
		config.Lg("chat_repo", "GetChatById").Error(err.Error())
		return domainChat.Chat{}, err
	}

	c.UserName = userName

	return c, nil
}


func (r *Repository) GetUserChats(userName string ) ([]domainChat.Chat, error) {
	tx, err := r.dbConn.Begin(context.Background())
	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	defer tx.Rollback(context.Background())

	records := 0
	err = tx.QueryRow(
		context.Background(),
		"SELECT COUNT(1) FROM users WHERE username = $1", userName).Scan(&records)

	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	if records == 0 {
		config.Lg("chat_repo", "GetUserChats").Error(errors.New("User doesn't exist"))
		return nil, errors.New("User doesn't exist")
	}


	chats, err := r.getUserChatsInternal(&tx, userName)
	if err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}


	if err := tx.Commit(context.Background()); err != nil {
		config.Lg("chat_repo", "GetUserChats").Error(err.Error())
		return nil, err
	}

	return chats, nil
}


func (r *Repository) getUserChatsInternal(tx *pgx.Tx, userName string) ([]domainChat.Chat, error) {
	rows, err := (*tx).Query(
		context.Background(),
		"SELECT c.id, c.creation_time, c.last_msg_id, c.last_msg_content, c.last_msg_username, " +
			"c.last_msg_time, u.username, u.avatar, uc.last_read_msg_id, uc.new_messages " +
			"FROM " +
			"chats c JOIN uu_chat uc " +
			"ON c.id = uc.chat_id " +
			"JOIN users u " +
			"ON u.id = uc.collocutor_id " +
			"WHERE uc.user_id = (SELECT id FROM users WHERE username = $1) " +
			"ORDER BY c.last_msg_time DESC, c.creation_time DESC ",
		userName)

	if err != nil {
		config.Lg("chat_repo", "GetUserChatsInt").Error(err.Error())
		return nil,  err
	}

	defer rows.Close()

	chats := []domainChat.Chat{}
	for rows.Next() {
		c := domainChat.Chat{}
		err := rows.Scan(&c.Id, &c.CreationTime, &c.LastMsgId, &c.LastMsgContent, &c.LastMsgUsername,
			&c.LastMsgTime, &c.CollocutorName, &c.CollocutorAvatarLink, &c.LastReadMsgId, &c.NewMessages)

		c.UserName = userName

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


func (r *Repository) MarkAllMessagesRead(chatId int, userName string) error {
	_, err := r.dbConn.Exec(
		context.Background(),
		"UPDATE uu_chat " +
			"SET new_messages = 0, " +
				"last_read_msg_id = chats.last_msg_id " +
			"FROM chats " +
			"WHERE chat_id = chats.id AND " +
				"user_id = (SELECT id FROM users WHERE username = $1) AND " +
				"chat_id = $2 ",
		userName, chatId)

	if err != nil {
		config.Lg("chat_repo", "GetChatById").Error(err.Error())
		return err
	}

	return nil
}
