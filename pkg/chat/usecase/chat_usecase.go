package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	proto "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Usecase struct {
	client  proto.ChatClient
	conn grpc.ClientConnInterface
}

func NewUsecase(conn grpc.ClientConnInterface) *Usecase {
	return &Usecase{
		conn: conn,
		client:  proto.NewChatClient(conn),
	}
}



func (uc *Usecase) CreateChat(req *domainChat.ChatCreateReq, userId int) (domainChat.ChatResp, error) {
	pResp, err := uc.client.CreateChat(context.Background(), &proto.ChatCreateReq{
		UserId: int32(userId),
		CollocutorName: req.CollocutorName,
	})

	if err != nil {
		uc.errorHandle(err)
		return domainChat.ChatResp{}, err
	}

	resp := domainChat.ChatResp{
		Id: int(pResp.Id),
		CreationTime: pResp.CreationTime.AsTime(),
		LastMsgContent: pResp.LastMsgContent,
		LastMsgUsername: pResp.LastMsgUsername,
		LastMsgTime: pResp.LastMsgTime.AsTime(),
		CollocutorName: pResp.CollocutorName,
		CollocutorAvatarLink: pResp.CollocutorAvatarLink,
		NewMessages: int(pResp.NewMessages),
	}

	return resp, nil
}

func (uc *Usecase) GetChatById(chatId int, userId int) (domainChat.ChatResp, error) {
	pResp, err := uc.client.GetChatById(context.Background(), &proto.GetChatByIdReq{
		UserId: int32(userId),
		ChatId: int32(chatId),
	})

	if err != nil {
		config.Lg("chat_usecase", "GetChatById").Error("client: ", err.Error())
		return domainChat.ChatResp{}, err
	}

	resp := domainChat.ChatResp{
		Id: int(pResp.Id),
		CreationTime: pResp.CreationTime.AsTime(),
		LastMsgContent: pResp.LastMsgContent,
		LastMsgUsername: pResp.LastMsgUsername,
		LastMsgTime: pResp.LastMsgTime.AsTime(),
		CollocutorName: pResp.CollocutorName,
		CollocutorAvatarLink: pResp.CollocutorAvatarLink,
		NewMessages: int(pResp.NewMessages),
	}

	return resp, nil
}


func (uc *Usecase) GetUserChats(userId int) ([]domainChat.ChatResp, error) {
	pRespArr, err := uc.client.GetUserChats(context.Background(), &proto.Id{
		Id: int32(userId),
	})

	if err != nil {
		config.Lg("chat_usecase", "GetUserChats").Error("client: ", err.Error())
		return nil, err
	}

	respArr := []domainChat.ChatResp{}
	for _, pResp := range pRespArr.Chats {
		respArr = append(respArr, domainChat.ChatResp{
			Id: int(pResp.Id),
			CreationTime: pResp.CreationTime.AsTime(),
			LastMsgContent: pResp.LastMsgContent,
			LastMsgUsername: pResp.LastMsgUsername,
			LastMsgTime: pResp.LastMsgTime.AsTime(),
			CollocutorName: pResp.CollocutorName,
			CollocutorAvatarLink: pResp.CollocutorAvatarLink,
			NewMessages: int(pResp.NewMessages),
		})
	}

	return respArr, nil
}

func (uc *Usecase) MarkAllMessagesRead(chatId int, userId int) error {
	_, err := uc.client.MarkMessagesRead(context.Background(), &proto.MarkMessagesReadReq{
		UserId: int32(userId),
		ChatId: int32(chatId),
	})

	if err != nil {
		config.Lg("chat_usecase", "MarkAllMessagesRead").Error("client: ", err.Error())
		return err
	}

	return nil
}

func (uc *Usecase) CreateMessage(mReq *domainChat.CreateMessageReq, userId int) (domainChat.MessageResp, error) {
	pResp, err := uc.client.CreateMessage(context.Background(), &proto.CreateMessageReq{
		UserId: int32(userId),
		ChatId: int32(mReq.ChatId),
		Content: mReq.Content,
	})

	if err != nil {
		config.Lg("chat_usecase", "MarkAllMessagesRead").Error("client: ", err.Error())
		return domainChat.MessageResp{}, err
	}

	resp := domainChat.MessageResp{
		Id: int(pResp.Id),
		Content: pResp.Content,
		CreationTime: pResp.CreationTime.AsTime(),
		ChatId: int(pResp.ChatId),
		UserName: pResp.UserName,
		UserAvatarLink: pResp.UserAvatarLink,
	}

	return resp, nil
}

func (uc *Usecase) DeleteMessage(msgId int) error {
	_, err := uc.client.DeleteMessage(context.Background(), &proto.Id{
		Id: int32(msgId),
	})

	if err != nil {
		config.Lg("chat_usecase", "DeleteMessage").Error("client: ", err.Error())
		return err
	}

	return nil
}

func (uc *Usecase) GetLastNMessages(mReq *domainChat.GetLastNMessagesReq) ([]domainChat.MessageResp, error) {
	pRespArr, err := uc.client.GetLastNMessages(context.Background(), &proto.GetLastNMessagesReq{
		ChatId: int32(mReq.ChatId),
		NMessages: int32(mReq.NMessages),
	})

	if err != nil {
		config.Lg("chat_usecase", "GetLastNMessages").Error("client: ", err.Error())
		return nil, err
	}

	respArr := []domainChat.MessageResp{}
	for _, pResp := range pRespArr.Messages {
		respArr = append(respArr, domainChat.MessageResp{
			Id: int(pResp.Id),
			Content: pResp.Content,
			CreationTime: pResp.CreationTime.AsTime(),
			ChatId: int(pResp.ChatId),
			UserName: pResp.UserName,
			UserAvatarLink: pResp.UserAvatarLink,
		})
	}

	return respArr, nil

}


func (uc *Usecase) GetNMessagesBefore(mReq *domainChat.GetNMessagesBeforeReq) ([]domainChat.MessageResp, error) {
	pRespArr, err := uc.client.GetNMessagesBefore(context.Background(), &proto.GetNMessagesReq{
		ChatId: int32(mReq.ChatId),
		NMessages: int32(mReq.NMessages),
		MessageId: int32(mReq.BeforeMessageId),
	})

	if err != nil {
		config.Lg("chat_usecase", "GetNMessagesBefore").Error("client: ", err.Error())
		return nil, err
	}

	respArr := []domainChat.MessageResp{}
	for _, pResp := range pRespArr.Messages {
		respArr = append(respArr, domainChat.MessageResp{
			Id: int(pResp.Id),
			Content: pResp.Content,
			CreationTime: pResp.CreationTime.AsTime(),
			ChatId: int(pResp.ChatId),
			UserName: pResp.UserName,
			UserAvatarLink: pResp.UserAvatarLink,
		})
	}

	return respArr, nil
}



func (uc *Usecase) errorHandle(err error) {
	errStatus, _ := status.FromError(err)
	config.Lg("chat_usecase", "errorHandle").Error(errStatus.Message())

	//switch errStatus.Code() {
	//case codes.Unavailable:
	//
	//}
}