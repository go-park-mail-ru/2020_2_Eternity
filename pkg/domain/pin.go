package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

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

func (p PinReq) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required, validation.Length(0, 250)),
		validation.Field(&p.Content, validation.Required, validation.Length(0, 500)),
	)
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
