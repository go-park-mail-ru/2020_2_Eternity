package domain

import "time"


const (
	NotificationKey = "notification"

	NoteCommentRespType = "NoteCommentResp"
	NotePinRespType = "NotePinResp"
	NoteFollowRespType = "NoteFollowResp"
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

// Notes for likes
// Notes for chat
