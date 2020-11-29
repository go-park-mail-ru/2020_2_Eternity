package domainChat

import "time"


// model
type Chat struct {
	Id int

	CreationTime time.Time
	LastMsgId int
	LastMsgContent string
	LastMsgUsername string
	LastMsgTime time.Time

	UserName string
	CollocutorName string
	CollocutorAvatarLink string
	LastReadMsgId int
	NewMessages int
}



type ChatCreateReq struct {
	CollocutorName string  `json:"collocutor_name"`
}



type MarkMessagesReadReq struct {
	UserName string  `json:"username"`
}




type ChatResp struct {
	Id int `json:"id"`
	CreationTime time.Time `json:"creation_time"`
	LastMsgContent string `json:"last_msg_content"`
	LastMsgUsername string `json:"last_msg_username"`
	LastMsgTime time.Time  `json:"last_msg_time"`

	CollocutorName string `json:"collocutor_name"`
	CollocutorAvatarLink string `json:"collocutor_ava"`
	NewMessages int `json:"new_messages"`
}



