package model

import (
	"sync"
)

type SignUpUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Age      int    `json:"age"`
}

type IUsers interface {
	CreateUser(User) error
}

type MockUsers struct {
	mu    *sync.Mutex
	users map[int]*User
}

func NewMockUsers() *MockUsers {
	return &MockUsers{
		mu:    &sync.Mutex{},
		users: make(map[int]*User, 0),
	}
}

func (u *MockUsers) CheckUser(user User) error {
	return nil
}

func (u *MockUsers) CreateUser(user User) error {
	if err := u.CheckUser(user); err != nil {
		return err
	}
	u.mu.Lock()
	id := len(u.users)
	u.users[id] = &User{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}

	u.mu.Unlock()
	return nil
}
