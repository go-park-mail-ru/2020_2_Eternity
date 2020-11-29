package ws

type Message struct {
	UserId int  `json:"-"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}
