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



func (uc *Usecase) CreateChat(req *domainChat.ChatCreateReq) (domainChat.ChatResp, error) {
	pResp, err := uc.client.CreateChat(context.Background(), &proto.ChatCreateReq{
		UserName: req.UserName,
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


func (uc *Usecase) errorHandle(err error) {
	errStatus, _ := status.FromError(err)
	config.Lg("chat_usecase", "errorHandle").Error(errStatus.Message())

	//switch errStatus.Code() {
	//case codes.Unavailable:
	//
	//}
}