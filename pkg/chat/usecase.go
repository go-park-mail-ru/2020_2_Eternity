package chat

import domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"

type IUsecase interface {
	CreateChat(req *domainChat.ChatCreateReq, userId int) (domainChat.ChatResp, error)
	GetChatById(chatId int, userId int) (domainChat.ChatResp, error)
	GetUserChats(userId int) ([]domainChat.ChatResp, error)
	MarkAllMessagesRead(chatId int, userId int) error
	CreateMessage(mReq *domainChat.CreateMessageReq, userId int) (domainChat.MessageResp, error)
	DeleteMessage(msgId int) error
	GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.MessageResp, error)
}


/*
	1) POST /chat - create chat

	type ChatCreateReq struct {
		CollocutorName string  `json:"collocutor_name"`
	}

	2) GET /chat/(chat_id) - get_chat by id

	3) GET /chat - get user chats

	4) PUT /chat - mark all messages read

	type MarkMessagesReadReq struct {
		ChatId int  `json:"chat_id"`
	}

 */