package ws

type MessageReq struct {
	UserId int  `json:"-"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}


type MessageResp struct {
	UserId int `json:"-"`
	Type string  `json:"type"`
	Status int  `json:"status"`
	Data []byte `json:"data"`
}