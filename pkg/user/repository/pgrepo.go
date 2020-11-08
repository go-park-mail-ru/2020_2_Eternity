package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
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

func (r *Repository) CreateUser(user *api.SignUp) (*domain.User, error) {
	u := &domain.User{}
	if _, err := r.dbConn.Exec(context.Background(), "insert into users(username, email, password, name, surname, description, birthdate, reg_date, avatar) values($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		user.Username, user.Email, user.Password, user.Name, user.Surname, user.Description, user.BirthDate, time.Now(), ""); err != nil {
		config.Lg("user", "CreateUser").Error("r.CreateUser: ", err.Error())
		return u, errors.New("user exists")
	}
	return u, nil
}

func (r *Repository) GetUser(id int) (*domain.User, error) {
	u := &domain.User{}
	row := r.dbConn.QueryRow(context.Background(), "select username, name, surname, description, password, email, birthdate, avatar from users where id=$1", id)
	if err := row.Scan(&u.Username, &u.Name, &u.Surname, &u.Description, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		config.Lg("user", "GetUser").Error("r.GetUser: ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) GetUserByName(username string) (*domain.User, error) {
	u := &domain.User{
		Username: username,
	}
	row := r.dbConn.QueryRow(context.Background(), "select id, password, email, birthdate, avatar from users where username=$1", username)
	if err := row.Scan(&u.ID, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		config.Lg("user", "GetUserByName").Error("r.GetUserByName: ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(id int, profile *api.UpdateUser) (*domain.User, error) {
	u := &domain.User{}

	if profile.Email != "" {
		if _, err := r.dbConn.Exec(context.Background(), "update users set email=$1 where id=$2",
			profile.Email, id); err != nil {
			config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
			return u, errors.New("fail update email")
		}
	}
	if profile.Username != "" {
		if _, err := r.dbConn.Exec(context.Background(), "update users set username=$1 where id=$2",
			profile.Username, id); err != nil {
			config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
			return u, errors.New("fail update username")
		}
		u.Username = profile.Username
	}

	if profile.Name != "" {
		if _, err := r.dbConn.Exec(context.Background(), "update users set name=$1 where id=$2",
			profile.Name, id); err != nil {
			config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
			return u, errors.New("fail update name")
		}
		u.Name = profile.Name
	}

	if profile.Surname != "" {
		if _, err := r.dbConn.Exec(context.Background(), "update users set surname=$1 where id=$2",
			profile.Surname, id); err != nil {
			config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
			return u, errors.New("fail update surname")
		}
		u.Surname = profile.Surname
	}

	if profile.Description != "" {
		if _, err := r.dbConn.Exec(context.Background(), "update users set description=$1 where id=$2",
			profile.Description, id); err != nil {
			config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
			return u, errors.New("fail update descr")
		}
		u.Description = profile.Description
	}
	return u, nil
}

func (r *Repository) UpdatePassword(id int, psswd string) error {
	if _, err := r.dbConn.Exec(context.Background(), "update users set password=$1 where id=$2", psswd, id); err != nil {
		config.Lg("user", "UpdatePassword").Error("r.UpdatePassword: ", err.Error())
		return errors.New("password error")
	}
	return nil
}

func (r *Repository) UpdateAvatar(id int, avatar string) error {
	if _, err := r.dbConn.Exec(context.Background(), "update users set avatar=$1 where id=$2", avatar, id); err != nil {
		config.Lg("user", "UpdateAvatar").Error("r.UpdateAvatar: ", err.Error())
		return errors.New("avatar doesnt update")
	}
	return nil
}

func (r *Repository) GetAvatar(id int) (error, string) {
	var avatar string
	row := r.dbConn.QueryRow(context.Background(), "select avatar from users where id=$1", id)
	if err := row.Scan(&avatar); err != nil {
		config.Lg("user", "GetAvatar").Error("r.GetAvatar: ", err.Error())
		return errors.New("user not found"), avatar
	}
	return nil, avatar
}

func (r *Repository) DeleteByName(username string) error {
	return nil
}

func (r *Repository) Follow(following int, id int) error {
	if _, err := r.dbConn.Exec(context.Background(), "insert into follows(id1, id2) values($1, $2)", following, id); err != nil {
		config.Lg("user", "Follow").Error("r.UpdatePassword: ", err.Error())
		return err
	}
	return nil
}

func (r *Repository) UnFollow(unfollowing int, id int) error {
	if _, err := r.dbConn.Exec(context.Background(), "delete from follows where id1=$1 and id2=$2", unfollowing, id); err != nil {
		config.Lg("user", "UnFollow").Error("r.UnFollow: ", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByNameWithFollowers(username string) (*domain.User, error) {
	u := &domain.User{
		Username: username,
	}
	row := r.dbConn.QueryRow(context.Background(), "select users.id, password, name, surname, description, avatar, followers, following from users join stats on users.id = stats.id where username=$1", username)
	if err := row.Scan(&u.ID, &u.Password, &u.Name, &u.Surname, &u.Description, &u.Avatar, &u.Followers, &u.Following); err != nil {
		config.Lg("user", "GetUserByNameWithFollowers").Error("r.GetUserByNameWithFollowers ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) GetFollowersIds(userId int) ([]int, error) {
	followers := []int{}
	row := r.dbConn.QueryRow(
		context.Background(),
		"select ARRAY(select id1 from follows where id2 = $1)",
		userId)
	if err := row.Scan(&followers); err != nil {
		config.Lg("user", "GetFollowersIds").Println(err)
		return []int{}, err
	}
	return followers, nil
}

func (r *Repository) GetFollowers(username string) ([]domain.User, error) {
	rows, err := r.dbConn.Query(context.Background(), "select us.username, us.avatar from (users as u join follows on u.id = id2) "+
		"p join users as us on p.id1 = us.id where p.username = $1", username)
	if err != nil {
		config.Lg("user", "GetFollowers").Error("r.GetFollowers ", err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		u := domain.User{}
		if err := rows.Scan(&u.Username, &u.Avatar); err != nil {
			config.Lg("user", "GetFollowers").Error("r.GetFollowers ", err.Error())
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func (r *Repository) GetFollowing(username string) ([]domain.User, error) {
	rows, err := r.dbConn.Query(context.Background(), "select us.username, us.avatar from (users as u join follows on u.id = id1) "+
		"p join users as us on p.id2 = us.id where p.username = $1", username)
	if err != nil {
		config.Lg("user", "GetFollowing").Error("r.GetFollowing ", err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		u := domain.User{}
		if err := rows.Scan(&u.Username, &u.Avatar); err != nil {
			config.Lg("user", "GetFollowing").Error("r.GetFollowing ", err.Error())
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
