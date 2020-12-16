package domain

type ReportReq struct {
	Message string `json:"message"`
	Type    int    `json:"type"`
	PinId   int    `json:"pin_id"`
}

type Report struct {
	Id       int    `json:"id"`
	Message  string `json:"message"`
	OwnerId  string `json:"owner_id"`
	Type     int    `json:"type"`
	PinId    string `json:"pin_id"`
	PinOwner string `json:"pin_owner"`
}
