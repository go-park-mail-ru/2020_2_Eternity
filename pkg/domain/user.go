package domain

import "time"

type User struct {
	ID        int       `json:"-"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	BirthDate time.Time `json:"date"`
	Avatar    string    `json:"avatar"`
	Followers int       `json:"followers"`
	Following int       `json:"following"`
}
