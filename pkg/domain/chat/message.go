package domainChat

import "time"

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
	UserName string
	ChatId int
	Content string
}

type GetLastNMessagesReq struct {
	UserName string
	ChatId int
	NMessages int
}

type GetNMessagesBeforeReq struct {
	UserName string
	ChatId int
	NMessages int
	BeforeMessageId int
}


type FileInfo struct {
	Id int
	Type string
	Name string
	Link string
}


type MessageResp struct {
	Id int
	Content string
	CreationTime time.Time
	ChatId int
	IsRead bool
	UserName string
	UserAvatarLink string
	Files []FileInfo
}
