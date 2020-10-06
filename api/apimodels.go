package api

import "time"

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

type CreatePinApi struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
