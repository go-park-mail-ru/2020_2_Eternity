package usecase

import (
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat"
)

type Usecase struct {
	repo  chat.IRepository
}

func NewUsecase(r chat.IRepository) *Usecase {
	return &Usecase{
		repo:  r,
	}
}

func (uc *Usecase) CreateChat(req *domainChat.ChatCreateReq) (domainChat.ChatResp, error) {
	ch := domainChat.Chat{}
	err := uc.repo.StoreChat(&ch, req.UserName, req.CollocutorName)

	resp := domainChat.ChatResp{
		Id: ch.Id,
		CreationTime: ch.CreationTime,
		LastMsgContent: ch.LastMsgContent,
		LastMsgUsername: ch.LastMsgUsername,
		LastMsgTime: ch.LastMsgTime,
		CollocutorName: ch.CollocutorName,
		CollocutorAvatarLink: ch.CollocutorAvatarLink,
		NewMessages: ch.NewMessages,
	}

	return resp, err
}


//GetChatById(chatId int, userName string) (domainChat.Chat, error)
//GetUserChats(userName string ) ([]domainChat.Chat, error)
//MarkAllMessagesRead(chatId int, userName string) error
//StoreMessage(mReq *domainChat.CreateMessageReq) (domainChat.Message, error)
//DeleteMessage(msgId int) error
//GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.Message, error)
