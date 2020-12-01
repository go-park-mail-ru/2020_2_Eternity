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

func (uc *Usecase) CreateChat(req *domainChat.ChatCreateReq, userId int) (domainChat.ChatResp, error) {
	ch := domainChat.Chat{}
	err := uc.repo.StoreChat(&ch, userId, req.CollocutorName)

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

func (uc *Usecase) GetChatById(chatId int, userId int) (domainChat.ChatResp, error) {
	ch, err := uc.repo.GetChatById(chatId, userId)
	if err != nil {
		return domainChat.ChatResp{}, err
	}

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

	return resp, nil
}

func (uc *Usecase) GetUserChats(userId int) ([]domainChat.ChatResp, error) {
	chats, err := uc.repo.GetUserChats(userId)
	if err != nil {
		return nil, err
	}

	resp := []domainChat.ChatResp{}
	for _, ch := range chats {
		resp = append(resp, domainChat.ChatResp{
			Id: ch.Id,
			CreationTime: ch.CreationTime,
			LastMsgContent: ch.LastMsgContent,
			LastMsgUsername: ch.LastMsgUsername,
			LastMsgTime: ch.LastMsgTime,
			CollocutorName: ch.CollocutorName,
			CollocutorAvatarLink: ch.CollocutorAvatarLink,
			NewMessages: ch.NewMessages,
		})
	}

	return resp, nil
}


func (uc *Usecase) MarkAllMessagesRead(chatId int, userId int) error {
	return uc.repo.MarkAllMessagesRead(chatId, userId)
}

func (uc *Usecase) CreateMessage(mReq *domainChat.CreateMessageReq, userId int) (domainChat.MessageResp, error) {
	m, err := uc.repo.StoreMessage(mReq, userId)
	if err != nil {
		return domainChat.MessageResp{}, err
	}

	resp := domainChat.MessageResp{
		Id: m.Id,
		Content: m.Content,
		CreationTime: m.CreationTime,
		ChatId: m.ChatId,
		UserName: m.UserName,
		UserAvatarLink: m.UserAvatarLink,
		Files: m.Files,
	}

	return resp, nil
}

func (uc *Usecase) DeleteMessage(msgId int) error {
	return uc.repo.DeleteMessage(msgId)
}


func (uc *Usecase) GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.MessageResp, error) {
	msgs, err := uc.repo.GetLastNMessages(mReq)
	if err != nil {
		return nil, err
	}

	resp := []domainChat.MessageResp{}
	for _, m := range msgs {
		resp = append(resp, domainChat.MessageResp{
			Id: m.Id,
			Content: m.Content,
			CreationTime: m.CreationTime,
			ChatId: m.ChatId,
			UserName: m.UserName,
			UserAvatarLink: m.UserAvatarLink,
			Files: m.Files,
		})
	}

	return resp, nil
}

func (uc *Usecase) GetNMessagesBefore(mReq *domainChat.GetNMessagesBeforeReq) ([]domainChat.MessageResp, error) {
	msgs, err := uc.repo.GetNMessagesBefore(mReq)
	if err != nil {
		return nil, err
	}

	resp := []domainChat.MessageResp{}
	for _, m := range msgs {
		resp = append(resp, domainChat.MessageResp{
			Id: m.Id,
			Content: m.Content,
			CreationTime: m.CreationTime,
			ChatId: m.ChatId,
			UserName: m.UserName,
			UserAvatarLink: m.UserAvatarLink,
			Files: m.Files,
		})
	}

	return resp, nil
}
