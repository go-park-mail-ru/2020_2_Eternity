package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Age      int    `json:"age"`
}

type IUsers interface {
	CreateUser(User) error
	CheckUser(User) error
}
