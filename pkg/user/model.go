package user

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"log"
	"time"
)

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

func (u *User) CreateUser() error {
	_, err := config.Db.Exec("insert into users(username, email, password, birthdate, reg_date, avatar) values($1, $2, $3, $4, $5, $6)",
		u.Username, u.Email, u.Password, u.BirthDate, time.Now(), "http://127.0.0.1:8008/images/avatar/default.jpeg")
	if err != nil {
		return errors.New("user exists")
	}
	return nil
}

func (u *User) GetUser() bool {
	row := config.Db.QueryRow("select username, password, email, birthdate, avatar from users where id=$1", u.ID)
	if err := row.Scan(&u.Username, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		return false
	}
	return true
}

func (u *User) GetUserByName() bool {
	row := config.Db.QueryRow("select id, password, email, birthdate, avatar from users where username=$1", u.Username)
	if err := row.Scan(&u.ID, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (u *User) UpdateUser(profile *api.UpdateUser) error {
	if _, err := config.Db.Exec("update users set username=$1, email=$2, birthdate=$3 where id=$4",
		profile.Username, profile.Email, profile.BirthDate, u.ID); err != nil {
		return errors.New("username or email exists")
	}
	u.Username = profile.Username
	u.Email = profile.Email
	u.BirthDate = profile.BirthDate
	return nil
}

func (u *User) UpdatePassword(psswd string) error {
	if _, err := config.Db.Exec("update users set password=$1 where id=$2", psswd, u.ID); err != nil {
		return errors.New("password error")
	}
	u.Password = psswd
	return nil
}

func (u *User) UpdateAvatar(avatar string) error {
	if _, err := config.Db.Exec("update users set avatar=$1 where id=$2", avatar, u.ID); err != nil {
		fmt.Println(err)
		return errors.New("avatar doesnt update")
	}
	u.Avatar = avatar
	return nil
}

func (u *User) GetAvatar() error {
	row := config.Db.QueryRow("select avatar from users where id=$1", u.ID)
	if err := row.Scan(&u.Avatar); err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (u *User) DeleteByName(username string) error {
	return nil
}

func (u *User) Follow(id int) error {
	// вставить строку, если не существует такая(todo: проверку на существование)
	if _, err := config.Db.Exec("insert into follows(id1, id2) values($1, $2)", u.ID, id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) UnFollow(id int) error {
	if _, err := config.Db.Exec("delete from follows where id1=$1 and id2=$2", u.ID, id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) GetUserByNameWithFollowers() error {
	row := config.Db.QueryRow("select users.id, avatar, followers, following from users join stats on users.id = stats.id where username=$1", u.Username)
	if err := row.Scan(&u.ID, &u.Avatar, &u.Followers, &u.Following); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
