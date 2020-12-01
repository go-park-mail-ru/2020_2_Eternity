package domain

import "time"


const (
	NotificationKey = "notification"

	NoteCommentRespType = "NoteCommentResp"
	NotePinRespType = "NotePinResp"
	NoteFollowRespType = "NoteFollowResp"
	NoteChatRespType = "NoteChatResp"
	NoteMessageRespType = "NoteMessageResp"
)

// Model
type Notification struct {
	Id           int       `json:"id"`
	ToUserId     int       `json:"to_id"`
	Type         int       `json:"type"`
	EncodedData  []byte    `json:"msg"`
	CreationTime time.Time `json:"time"`
	IsRead 		 bool      `json:"is_read"`
}

// Delivery

type NoteResp struct {
	Id           int       `json:"id"`
	Type         int       `json:"type"`
	EncodedData  []byte    `json:"data"`
	CreationTime time.Time `json:"creation_time"`
	IsRead 		 bool      `json:"is_read"`
}


// Notification types

type NotePin struct {
	Id      int    `bson:"id"`
	Title   string `bson:"title"`
	ImgLink string `bson:"img_link"`
	UserId  int    `bson:"user_id"`
}

type NoteComment struct {
	Id      int 	`bson:"id"`
	Path    []int32	`bson:"path"`
	Content string  `bson:"content"`
	PinId   int     `bson:"pin_id"`
	UserId  int     `bson:"userid"`
}


type NoteFollow struct {
	FollowerId  int `bson:"follower_id"`
	UserId  int `bson:"user_id"`
}


type NoteChat struct {
	Id int `json:"id"`
	CreatorId int `json:"-"`
	CreationTime time.Time `json:"creation_time"`
	LastMsgContent string `json:"last_msg_content"`
	LastMsgUsername string `json:"last_msg_username"`
	LastMsgTime time.Time  `json:"last_msg_time"`

	CollocutorName string `json:"collocutor_name"`
	CollocutorAvatarLink string `json:"collocutor_ava"`
	NewMessages int `json:"new_messages"`
}

type NoteMessage struct {
	Id int  `json:"id"`
	CreatorId int `json:"-"`
	Content string  `json:"content"`
	CreationTime time.Time  `json:"time"`
	ChatId int `json:"chat_id"`
	UserName string  `json:"username"`
	UserAvatarLink string  `json:"avatar"`
}

// Notes for likes

