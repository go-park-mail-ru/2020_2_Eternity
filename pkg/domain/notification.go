package domain

type Notification struct {
	FromUserId int `json:"from_id"`
	ToUserId int `json:"to_id"`
	Type int `json:"type"`
	Msg string `json:"msg"`
}
