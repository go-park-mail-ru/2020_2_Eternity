package domainChat

import "time"

var (
	CreateMessageReqType = "CreateMessageReq"
	DeleteMessageReqType = "DeleteMessageReq"
	GetLastNMessagesReqType = "GetLastNMessagesReq"
	GetNMessagesBeforeReqType = "GetNMessagesBeforeReq"


	CreateMessageRespType = "CreateMessageResp"
	DeleteMessageRespType = "DeleteMessageResp"
	GetLastNMessagesRespType = "GetLastNMessagesResp"
	GetNMessagesBeforeRespType = "GetNMessagesBeforeResp"
)

// model
type Message struct {
	Id int
	Content string
	CreationTime time.Time
	ChatId int
	UserId int
	UserName string
	UserAvatarLink string
	Files []FileInfo
}




type CreateMessageReq struct {
	ChatId int  `json:"chat_id"`
	Content string  `json:"content"`
}

type DeleteMessageReq struct {
	MsgId int `json:"message_id"`
}

type GetLastNMessagesReq struct {
	ChatId int	`json:"chat_id"`
	NMessages int `json:"messages"`
}

type GetNMessagesBeforeReq struct {
	ChatId int  `json:"chat_id"`
	NMessages int `json:"n_messages"`
	BeforeMessageId int  `json:"message_id"`
}


type FileInfo struct {
	Id int
	Type string
	Name string
	Link string
}


type MessageResp struct {
	Id int  `json:"id"`
	Content string  `json:"content"`
	CreationTime time.Time  `json:"time"`
	ChatId int `json:"chat_id"`
	UserName string  `json:"username"`
	UserAvatarLink string  `json:"avatar"`
	Files []FileInfo  `json:"-"`
}
