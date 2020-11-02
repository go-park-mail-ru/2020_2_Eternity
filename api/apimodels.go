package api

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"time"
)

type GetProfile struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"date"`
}

type SignUp struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"date"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"date"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type Follow struct {
	Username string `json:"username"`
}

type UserPage struct {
	Username  string           `json:"username"`
	Followers int              `json:"followers"`
	Following int              `json:"following"`
	PinsList  []domain.PinResp `json:"pins_list"`
}
