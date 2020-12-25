package repository

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/lib/pq"
	"strconv"
	"strings"
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
	if err := r.dbConn.QueryRow("insert into users(username, email, password, name, surname, description, birthdate, reg_date, avatar) values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id",
		user.Username, user.Email, user.Password, user.Name, user.Surname, user.Description, user.BirthDate, time.Now(), "").Scan(&u.ID); err != nil {
		config.Lg("user", "CreateUser").Error("r.CreateUser: ", err.Error())
		return u, errors.New("user exists")
	}
	return u, nil
}

func (r *Repository) GetUser(id int) (*domain.User, error) {
	u := &domain.User{}
	row := r.dbConn.QueryRow("select username, name, surname, description, password, email, birthdate, avatar from users where id=$1", id)
	if err := row.Scan(&u.Username, &u.Name, &u.Surname, &u.Description, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		config.Lg("user", "GetUser").Error("r.GetUser: ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) GetUserByName(username string) (*domain.User, error) {
	u := &domain.User{}
	row := r.dbConn.QueryRow("select id, username, password, email, birthdate, avatar from users where lower(username)=lower($1)", username)
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.BirthDate, &u.Avatar); err != nil {
		config.Lg("user", "GetUserByName").Error("r.GetUserByName: ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(id int, profile *api.UpdateUser) (*domain.User, error) {
	u := &domain.User{}

	i := 0
	query := "update users set "
	var queryParams []string
	var values []interface{}
	if profile.Email != "" {
		i++
		queryParams = append(queryParams, "email=$"+strconv.Itoa(i))
		values = append(values, profile.Email)
		u.Email = profile.Email
	}
	if profile.Username != "" {
		i++
		queryParams = append(queryParams, "username=$"+strconv.Itoa(i))
		values = append(values, profile.Username)
		u.Username = profile.Username
	}

	if profile.Name != "" {
		i++
		queryParams = append(queryParams, "name=$"+strconv.Itoa(i))
		values = append(values, profile.Name)
		u.Name = profile.Name
	}

	if profile.Surname != "" {
		i++
		queryParams = append(queryParams, "surname=$"+strconv.Itoa(i))
		values = append(values, profile.Surname)
		u.Surname = profile.Surname
	}

	if profile.Description != "" {
		i++
		queryParams = append(queryParams, "description=$"+strconv.Itoa(i))
		values = append(values, profile.Description)
		u.Description = profile.Description
	}
	query += strings.Join(queryParams, ",")
	i++
	query += " where id=$" + strconv.Itoa(i)
	values = append(values, id)
	if _, err := r.dbConn.Exec(query,
		values...); err != nil {
		config.Lg("user", "UpdateUser").Error("r.UpdateUser: ", err.Error())
		return u, errors.New("fail update descr")
	}
	return u, nil
}

func (r *Repository) UpdatePassword(id int, psswd string) error {
	if _, err := r.dbConn.Exec("update users set password=$1 where id=$2", psswd, id); err != nil {
		config.Lg("user", "UpdatePassword").Error("r.UpdatePassword: ", err.Error())
		return errors.New("password error")
	}
	return nil
}

func (r *Repository) UpdateAvatar(id int, avatar string) error {
	if _, err := r.dbConn.Exec("update users set avatar=$1 where id=$2", avatar, id); err != nil {
		config.Lg("user", "UpdateAvatar").Error("r.UpdateAvatar: ", err.Error())
		return errors.New("avatar doesnt update")
	}
	return nil
}

func (r *Repository) GetAvatar(id int) (error, string) {
	var avatar string
	row := r.dbConn.QueryRow("select avatar from users where id=$1", id)
	if err := row.Scan(&avatar); err != nil {
		config.Lg("user", "GetAvatar").Error("r.GetAvatar: ", err.Error())
		return errors.New("user not found"), avatar
	}
	return nil, avatar
}

func (r *Repository) Follow(following int, id int) (string, error) {
	row := r.dbConn.QueryRow("insert into follows(id1, id2) values($1, $2) "+
		"returning (select username from users where id = $1)", following, id)
	var username string
	if err := row.Scan(&username); err != nil {
		config.Lg("user", "Follow").Error("r.UpdatePassword: ", err.Error())
		return username, err
	}
	return username, nil
}

func (r *Repository) UnFollow(unfollowing int, id int) error {
	if _, err := r.dbConn.Exec("delete from follows where id1=$1 and id2=$2", unfollowing, id); err != nil {
		config.Lg("user", "UnFollow").Error("r.UnFollow: ", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByNameWithFollowers(username string) (*domain.User, error) {
	u := &domain.User{}
	row := r.dbConn.QueryRow("select users.id, username, password, name, surname, description, avatar, followers, following from users join stats on users.id = stats.id where lower(username)=lower($1)", username)
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.Surname, &u.Description, &u.Avatar, &u.Followers, &u.Following); err != nil {
		config.Lg("user", "GetUserByNameWithFollowers").Error("r.GetUserByNameWithFollowers ", err.Error())
		return u, err
	}
	return u, nil
}

func (r *Repository) GetFollowersIds(userId int) ([]int, error) {
	followersSql := []sql.NullInt32{}
	row := r.dbConn.QueryRow(
		"select ARRAY(select id1 from follows where id2 = $1)",
		userId)
	if err := row.Scan(pq.Array(&followersSql)); err != nil {
		config.Lg("user", "GetFollowersIds").Println(err)
		return []int{}, err
	}
	followers := make([]int, 0, len(followersSql))
	for _, f := range followersSql {
		followers = append(followers, int(f.Int32))
	}
	return followers, nil
}

func (r *Repository) GetFollowers(username string) ([]domain.User, error) {
	rows, err := r.dbConn.Query("select us.username, us.avatar from (users as u join follows on u.id = id2) "+
		"p join users as us on p.id1 = us.id where lower(p.username) = lower($1)", username)
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
	rows, err := r.dbConn.Query("select us.username, us.avatar from (users as u join follows on u.id = id1) "+
		"p join users as us on p.id2 = us.id where lower(p.username) = lower($1)", username)
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

func (r *Repository) IsFollowing(id int, username string) error {
	var u int
	if err := r.dbConn.QueryRow("select id2 from follows join users on users.id = follows.id2 where lower(username)=lower($1) and id1=$2", username, id).Scan(&u); err != nil {
		config.Lg("user", "IsFollowing").Error(err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetPopularUsers(limit int) ([]domain.UserSearch, error) {
	rows, err := r.dbConn.Query("select users.id, username, avatar, followers from users join "+
		"stats on users.id = stats.id order by followers desc limit $1", limit)

	if err != nil {
		config.Lg("user", "GetPopularUsers").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	users := make([]domain.UserSearch, 0, limit)
	u := domain.UserSearch{}
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Username, &u.Avatar, &u.Followers); err != nil {
			config.Lg("user", "GetPopularUsersScan").Error(err.Error())
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
