package domain

// Model
type Pin struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PictureName string `json:"picture_name"`
	UserId      int    `json:"user_id"`
}

// Create req
type PinReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Edit req
type PinEditReq struct {
	//Title   string `json:"title"`
	//Content string `json:"content"`
	// TODO (Paul S) Something here
}

type PinResp struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	ImgLink string `json:"img_link"`
	UserId  int    `json:"user_id"`
}
