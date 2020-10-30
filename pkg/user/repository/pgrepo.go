package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/models"
	"log"
	"time"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}

func (r *Repository) CreateUser(user *api.SignUp) (*models.User, error) {
	u := &models.User{}
	if _, err := r.dbConn.Exec(context.Background(), "insert into users(username, email, password, birthdate, reg_date, avatar) values($1, $2, $3, $4, $5, $6)",
		user.Username, user.Email, user.Password, user.BirthDate, time.Now(), "http://127.0.0.1:8008/images/avatar/default.jpeg"); err != nil {
		return u, errors.New("user exists")
	}
	return u, nil
}

func (r *Repository) GetUser(id int) (*models.User, error) {
	u := &models.User{}
	row := r.dbConn.QueryRow(context.Background(), "select username, password, email, birthdate, avatar from users where id=$1", id)
	if err := row.Scan(&u.Username, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		return u, err
	}
	return u, nil
}

func (r *Repository) GetUserByName(username string) (*models.User, error) {
	u := &models.User{
		Username: username,
	}
	row := r.dbConn.QueryRow(context.Background(), "select id, password, email, birthdate, avatar from users where username=$1", username)
	if err := row.Scan(&u.ID, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		fmt.Println(err)
		return u, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(id int, profile *api.UpdateUser) (*models.User, error) {
	u := &models.User{}
	if _, err := r.dbConn.Exec(context.Background(), "update users set username=$1, email=$2, birthdate=$3 where id=$4",
		profile.Username, profile.Email, profile.BirthDate, id); err != nil {
		return u, errors.New("username or email exists")
	}
	u.Username = profile.Username
	u.Email = profile.Email
	u.BirthDate = profile.BirthDate
	return u, nil
}

func (r *Repository) UpdatePassword(id int, psswd string) error {
	if _, err := r.dbConn.Exec(context.Background(), "update users set password=$1 where id=$2", psswd, id); err != nil {
		log.Println(psswd)
		return errors.New("password error")
	}
	log.Println(psswd)
	return nil
}

func (r *Repository) UpdateAvatar(id int, avatar string) error {
	if _, err := r.dbConn.Exec(context.Background(), "update users set avatar=$1 where id=$2", avatar, id); err != nil {
		fmt.Println(err)
		return errors.New("avatar doesnt update")
	}
	return nil
}

func (r *Repository) GetAvatar(id int) (error, string) {
	var avatar string
	row := r.dbConn.QueryRow(context.Background(), "select avatar from users where id=$1", id)
	if err := row.Scan(&avatar); err != nil {
		return errors.New("user not found"), avatar
	}
	return nil, avatar
}

func (r *Repository) DeleteByName(username string) error {
	return nil
}

func (r *Repository) Follow(following int, id int) error {
	if _, err := r.dbConn.Exec(context.Background(), "insert into follows(id1, id2) values($1, $2)", following, id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repository) UnFollow(unfollowing int, id int) error {
	if _, err := r.dbConn.Exec(context.Background(), "delete from follows where id1=$1 and id2=$2", unfollowing, id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repository) GetUserByNameWithFollowers(username string) (*models.User, error) {
	u := &models.User{
		Username: username,
	}

	row := r.dbConn.QueryRow(context.Background(), "select users.id, avatar, followers, following from users join stats on users.id = stats.id where username=$1", username)
	if err := row.Scan(&u.ID, &u.Avatar, &u.Followers, &u.Following); err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}
