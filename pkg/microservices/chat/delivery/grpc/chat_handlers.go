package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat"
	proto "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/chat"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		CollocutorName: pReq.CollocutorName,
	}

	resp, err := c.uc.CreateChat(&req, int(pReq.UserId))

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


func (c *ChatServer) GetChatById(ctx context.Context, pReq *proto.GetChatByIdReq) (*proto.ChatResp, error) {
	resp, err := c.uc.GetChatById(int(pReq.ChatId), int(pReq.UserId))
	if err != nil {
		config.Lg("chat_grpc_ms", "GetChatById").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	crTime, err := ptypes.TimestampProto(resp.CreationTime)
	if err != nil {
		config.Lg("chat_grpc_ms", "GetChatById").Error("Time ", err.Error())
	}

	lastMsgTime, err := ptypes.TimestampProto(resp.LastMsgTime)
	if err != nil {
		config.Lg("chat_grpc_ms", "GetChatById").Error("Time ", err.Error())
	}

	pResp := proto.ChatResp{
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

func (c *ChatServer) GetUserChats(ctx context.Context, id *proto.Id) (*proto.ChatRespArray, error) {
	respArr, err := c.uc.GetUserChats(int(id.Id))
	if err != nil {
		config.Lg("chat_grpc_ms", "GetUserChats").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	pReps := proto.ChatRespArray{}

	for _, resp := range respArr {
		crTime, err := ptypes.TimestampProto(resp.CreationTime)
		if err != nil {
			config.Lg("chat_grpc_ms", "GetUserChats").Error("Time ", err.Error())
		}

		lastMsgTime, err := ptypes.TimestampProto(resp.LastMsgTime)
		if err != nil {
			config.Lg("chat_grpc_ms", "GetUserChats").Error("Time ", err.Error())
		}

		pReps.Chats = append(pReps.Chats, &proto.ChatResp{
			Id: int32(resp.Id),
			CreationTime: crTime,
			LastMsgContent: resp.LastMsgContent,
			LastMsgUsername: resp.LastMsgUsername,
			LastMsgTime: lastMsgTime,
			CollocutorName: resp.CollocutorName,
			CollocutorAvatarLink: resp.CollocutorAvatarLink,
			NewMessages: int32(resp.NewMessages),
		})
	}

	return &pReps, nil
}


func (c *ChatServer) MarkMessagesRead(ctx context.Context, pReq *proto.MarkMessagesReadReq) (*empty.Empty, error) {
	if err := c.uc.MarkAllMessagesRead(int(pReq.ChatId), int(pReq.UserId)); err != nil {
		config.Lg("chat_grpc_ms", "MarkMessagesRead").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}


func (c *ChatServer) CreateMessage(ctx context.Context, pReq *proto.CreateMessageReq) (*proto.MessageResp, error) {
	req := domainChat.CreateMessageReq{
		ChatId: int(pReq.ChatId),
		Content: pReq.Content,
	}

	resp, err := c.uc.CreateMessage(&req, int(pReq.UserId))
	if  err != nil {
		config.Lg("chat_grpc_ms", "CreateMessage").Error("Usecase ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	crTime, err := ptypes.TimestampProto(resp.CreationTime)
	if err != nil {
		config.Lg("chat_grpc_ms", "CreateMessage").Error("Time ", err.Error())
	}

	pResp := proto.MessageResp{
		Id: int32(resp.Id),
		Content: resp.Content,
		CreationTime: crTime,
		ChatId: int32(resp.ChatId),
		UserName: resp.UserName,
		UserAvatarLink: resp.UserAvatarLink,
	}

	return &pResp, nil
}


func (c *ChatServer) DeleteMessage(ctx context.Context, id *proto.Id) (*empty.Empty, error) {
	if err := c.uc.DeleteMessage(int(id.Id)); err != nil {
		config.Lg("chat_grpc_ms", "DeleteMessage").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}


func (c *ChatServer) GetLastNMessages(ctx context.Context, pReq *proto.GetLastNMessagesReq) (*proto.MessageRespArray, error) {
	req := domainChat.GetLastNMessagesReq{
		ChatId: int(pReq.ChatId),
		NMessages: int(pReq.NMessages),
	}

	respArr, err := c.uc.GetLastNMessages(&req)
	if err != nil {
		config.Lg("chat_grpc_ms", "GetLastNMessages").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	pReps := proto.MessageRespArray{}

	for _, resp := range respArr {
		crTime, err := ptypes.TimestampProto(resp.CreationTime)
		if err != nil {
			config.Lg("chat_grpc_ms", "GetLastNMessages").Error("Time ", err.Error())
		}

		pReps.Messages = append(pReps.Messages, &proto.MessageResp{
			Id: int32(resp.Id),
			Content: resp.Content,
			CreationTime: crTime,
			ChatId: int32(resp.ChatId),
			UserName: resp.UserName,
			UserAvatarLink: resp.UserAvatarLink,
		})
	}

	return &pReps, nil
}
func (c *ChatServer) GetNMessagesBefore(ctx context.Context, pReq *proto.GetNMessagesReq) (*proto.MessageRespArray, error) {
	req := domainChat.GetNMessagesBeforeReq{
		ChatId: int(pReq.ChatId),
		NMessages: int(pReq.NMessages),
		BeforeMessageId: int(pReq.MessageId),
	}

	respArr, err := c.uc.GetNMessagesBefore(&req)
	if err != nil {
		config.Lg("chat_grpc_ms", "GetNMessagesBefore").Error("Usecase: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	pReps := proto.MessageRespArray{}

	for _, resp := range respArr {
		crTime, err := ptypes.TimestampProto(resp.CreationTime)
		if err != nil {
			config.Lg("chat_grpc_ms", "GetNMessagesBefore").Error("Time ", err.Error())
		}

		pReps.Messages = append(pReps.Messages, &proto.MessageResp{
			Id: int32(resp.Id),
			Content: resp.Content,
			CreationTime: crTime,
			ChatId: int32(resp.ChatId),
			UserName: resp.UserName,
			UserAvatarLink: resp.UserAvatarLink,
		})
	}

	return &pReps, nil
}




