package chat

import domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"

type IRepository interface {
	StoreChat(ch *domainChat.Chat, userId int, collocutorName string) error
	GetChatById(chatId int, userId int) (domainChat.Chat, error)
	GetUserChats(userId int) ([]domainChat.Chat, error)
	MarkAllMessagesRead(chatId int, userId int) error
	StoreMessage(mReq *domainChat.CreateMessageReq, userId int) (domainChat.Message, error)
	DeleteMessage(msgId int) error
	GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.Message, error)
}