package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Model
type Comment struct {
	Id       int
	Path     []int32
	Content  string
	PinId    int
	UserId   int
	Username string
}

type CommentCreateReq struct {
	IsRoot   bool   `json:"is_root"` // NOTE (Pavel S) if true, ParentId is not checked
	ParentId int    `json:"parent_id"`
	Content  string `json:"content"`
	PinId    int    `json:"pin_id"`
}

func (p CommentCreateReq) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Content, validation.Required, validation.Length(0, 200)),
	)
}

// TODO (Pavel S) CommentEditReq

type CommentResp struct {
	Id       int     `json:"id"`
	Path     []int32 `json:"path"`
	Content  string  `json:"content"`
	PinId    int     `json:"pin_id"`
	UserId   int     `json:"user_id"`
	Username string  `json:"username"`
}
