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
	Avatar string
	Files []FileInfo
}


type CreateMessageReq struct {
	UserName string
	ChatId int
	Content string
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
	UserName string
	Avatar string
	Files []FileInfo
}
