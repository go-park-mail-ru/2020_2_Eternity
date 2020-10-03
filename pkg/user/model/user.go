package model

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	BirthDate int    `json:"date"`
}

type IUsers interface {
	CreateUser(User) error
	CheckUser(string, string) (int, bool)
}
