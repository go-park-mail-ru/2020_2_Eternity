package model

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

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

func (u *User) CreateUser() error {
	_, err := config.Db.Exec("insert into users(username, email, password, age, reg_date) values($1, $2, $3, $4, $5)",
		u.Username, u.Email, u.Password, u.BirthDate, time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) CheckUser() (int, bool) {
	row := config.Db.QueryRow("select id, password from users where username=$1", u.Username)
	var id int
	var hash string
	if err := row.Scan(&id, &hash); err != nil {
		fmt.Println(err)
		return -1, false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password)); err != nil {
		fmt.Println(hash, u.Password)
		return -1, false
	}
	return id, true
}
