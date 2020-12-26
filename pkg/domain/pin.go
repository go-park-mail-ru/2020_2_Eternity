package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Model
type Pin struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	PictureName   string `json:"picture_name"`
	PictureHeight int    `json:"height,omitempty"`
	PictureWidth  int    `json:"width,omitempty"`
	UserId        int    `json:"user_id"`
	Username      string `json:"username,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
}

// Create req
type PinReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p PinReq) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Length(0, 250)),
		validation.Field(&p.Content, validation.Length(0, 500)),
	)
}

// Edit req
type PinEditReq struct {
	//Title   string `json:"title"`
	//Content string `json:"content"`
	// TODO (Paul S) Something here
}

type PinResp struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	PictureHeight int    `json:"height,omitempty"`
	PictureWidth  int    `json:"width,omitempty"`
	ImgLink       string `json:"img_link"`
	UserId        int    `json:"user_id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar,omitempty"`
}
