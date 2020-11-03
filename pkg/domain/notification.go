package domain

import "time"

// Model
type Notification struct {
	Id           int       `json:"id"`
	FromUserId   int       `json:"from_id"`
	ToUserId     int       `json:"to_id"`
	Type         int       `json:"type"`
	EncodedData  []byte    `json:"msg"`
	CreationTime time.Time `json:"time"`
}
