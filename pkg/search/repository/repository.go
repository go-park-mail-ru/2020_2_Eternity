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
	query := "select id, username, avatar from users where lower(username) like lower('%' || $1 || '%') " +
		"or lower(name || surname) like lower('%' || $1 || '%') "
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, username)
	if last > 0 {
		i++
		query += "and users.id < $ " + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += "order by users.id desc limit 15"
	rows, err := r.db.Query(context.Background(), query, placeholders...)
	if err != nil {
		config.Lg("search", "UserNameSearch").Error(err.Error())
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
	query := "select p.id, title, content, name, user_id, ts_rank(\"vec\", plainto_tsquery($1)) from (pins " +
		"left join pins_vectors on pins.id = pins_vectors.idv) p join pin_images on p.id = pin_images.pin_id " +
		"where \"vec\" @@ plainto_tsquery($1)"
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, title)
	if last > 0 {
		i++
		query += " and pins.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by ts_rank(\"vec\", plainto_tsquery($1)) desc limit 15"
	rows, err := r.db.Query(context.Background(), query, placeholders...)
	if err != nil {
		config.Lg("search", "GetPinsSearch").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var pins []domain.Pin
	for rows.Next() {
		pin := domain.Pin{}
		var rank float32
		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId, &rank); err != nil {
			config.Lg("search", "GetPinSearch").Error(err.Error())
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}
