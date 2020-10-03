package config

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (db *Database) CreateUser(user model.User) error {
	_, err := db.database.Exec("insert into users(username, email, password, age, reg_date) values($1, $2, $3, $4, $5)",
		user.Username, user.Email, user.Password, user.BirthDate, time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db *Database) CheckUser(username string, password string) (int, bool) {
	row := db.database.QueryRow("select id, password from users where username=$1", username)
	var id int
	var hash string
	if err := row.Scan(&id, &hash); err != nil {
		fmt.Println(err)
		return -1, false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		fmt.Println(hash, password)
		return -1, false
	}
	return id, true
}
