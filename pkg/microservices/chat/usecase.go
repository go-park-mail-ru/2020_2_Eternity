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