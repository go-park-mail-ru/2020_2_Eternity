package domainChat

import "time"

var (
	CreateMessageReqType = "CreateMessageReq"
	GetLastNMessagesReqType = "GetLastNMessagesReq"
	GetNMessagesBeforeReqType = "GetNMessagesBeforeReq"

	// ??
	MessageRespOne = "MessageRespOne"
	MessageRespMany = "MessageRespMany"
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
