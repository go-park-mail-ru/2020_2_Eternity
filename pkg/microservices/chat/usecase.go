package chat

import domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"

type IUsecase interface {
	CreateChat(req *domainChat.ChatCreateReq, userId int) (domainChat.ChatResp, error)
}