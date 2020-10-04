package model

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	BirthDate time.Time `json:"date"`
}

func (u *User) CreateUser() error {
	_, err := config.Db.Exec("insert into users(username, email, password, birthdate, reg_date) values($1, $2, $3, $4, $5)",
		u.Username, u.Email, u.Password, u.BirthDate, time.Now())
	if err != nil {
		return errors.New("user exists")
	}
	return nil
}

func (u *User) GetUser() bool {
	row := config.Db.QueryRow("select id, password, email, birthdate from users where username=$1", u.Username)
	if err := row.Scan(&u.ID, &u.Password, &u.Email, &u.BirthDate); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
