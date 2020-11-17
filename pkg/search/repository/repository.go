package repository

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"strconv"
)

type Repository struct {
	db database.IDbConn
}

func NewRepository(db database.IDbConn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUsersByName(username string, last int) ([]domain.UserSearch, error) {
	query := "select id, username, avatar from users where username like $1 "
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, username)
	if last > 0 {
		i++
		query += "and users.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by users.id desc limit 15"
	rows, err := r.db.Query(context.Background(), query, placeholders...)
	if err != nil {
		config.Lg("feed", "GetFeed").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []domain.UserSearch
	for rows.Next() {
		u := domain.UserSearch{}
		if err := rows.Scan(&u.ID, &u.Username, &u.Avatar); err != nil {
			config.Lg("search", "users").Error(err.Error())
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repository) GetPinsByTitle(title string, last int) ([]domain.Pin, error) {
	return nil, nil
}
