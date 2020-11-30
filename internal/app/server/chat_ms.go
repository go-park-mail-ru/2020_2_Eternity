package server

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"google.golang.org/grpc"
	"time"
)

func NewChatMsConnection() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())


	serverAddr := fmt.Sprintf("%s:%s", config.Conf.Web.Chat.Address, config.Conf.Web.Chat.Port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, opts...)

	if err != nil {
		config.Lg("server", "NewNewChatMsConnection").
			Fatalf("Fail to dial (ChatService): %v", err)
		return nil
	}

	return conn
}
