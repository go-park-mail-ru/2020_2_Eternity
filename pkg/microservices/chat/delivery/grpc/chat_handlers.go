package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/chat"
)


type ChatServer struct {
	uc chat.IUsecase
}

func NewChatServer(uc chat.IUsecase) *ChatServer{
	return &ChatServer{
		uc: uc,
	}
}

func (c *ChatServer) CreateChat(ctx context.Context, pReq *proto.ChatCreateReq) (*proto.ChatResp, error) {
	req := domainChat.ChatCreateReq{
		UserName: pReq.UserName,
		CollocutorName: pReq.CollocutorName,
	}

	resp, err := c.uc.CreateChat(&req)

	if err != nil {
		config.Lg("chat_grpc_ms", "CreateChat").Error("Usecase ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	crTime, err := ptypes.TimestampProto(resp.CreationTime)
	if err != nil {
		config.Lg("chat_grpc_ms", "CreateChat").Error("Time ", err.Error())
	}

	lastMsgTime, err := ptypes.TimestampProto(resp.LastMsgTime)
	if err != nil {
		config.Lg("chat_grpc_ms", "CreateChat").Error("Time ", err.Error())
	}

	pResp := proto.ChatResp {
		Id: int32(resp.Id),
		CreationTime: crTime,
		LastMsgContent: resp.LastMsgContent,
		LastMsgUsername: resp.LastMsgUsername,
		LastMsgTime: lastMsgTime,
		CollocutorName: resp.CollocutorName,
		CollocutorAvatarLink: resp.CollocutorAvatarLink,
		NewMessages: int32(resp.NewMessages),
	}

	return &pResp, nil
}


func (c *ChatServer) MarkMessagesRead(context.Context, *proto.MarkMessagesReadReq) (*proto.ChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkMessagesRead not implemented")
}


func (c *ChatServer) RouteWs(proto.Chat_RouteWsServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteWs not implemented")
}

