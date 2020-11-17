package domain

import "time"

type User struct {
	ID          int       `json:"-"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Description string    `json:"description"`
	BirthDate   time.Time `json:"date"`
	Avatar      string    `json:"avatar"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
}

type UserSearch struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
