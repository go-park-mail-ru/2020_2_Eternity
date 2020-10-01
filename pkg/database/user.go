package database

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"github.com/jackc/pgx"
	"time"
)

func (db *DB) CreateUser(user model.User) error {
	_, err := db.conn.Exec("insert into users(username, email, password, age, reg_date) values($1, $2, $3, $4, $5)",
		user.Username, user.Email, user.Password, user.Age, time.Now())
	if err, ok := err.(pgx.PgError); ok {
		return err
	}
	return nil
}

func (db *DB) CheckUser(user model.User) error {
	return nil
}
