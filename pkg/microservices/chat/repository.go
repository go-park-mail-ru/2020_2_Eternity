package chat

import domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"

type IRepository interface {
	StoreChat(ch *domainChat.Chat, userName string, collocutorName string) error
	GetChatById(chatId int, userName string) (domainChat.Chat, error)
	GetUserChats(userName string ) ([]domainChat.Chat, error)
	MarkAllMessagesRead(chatId int, userName string) error
	StoreMessage(mReq *domainChat.CreateMessageReq) (domainChat.Message, error)
	DeleteMessage(msgId int) error
	GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.Message, error)
}