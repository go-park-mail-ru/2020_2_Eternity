package repository

import (
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
		"or lower(name || surname) like lower('%' || $1 || '%')"
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, username)
	if last > 0 {
		i++
		query += " and users.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by users.id desc limit 15"
	rows, err := r.db.Query(query, placeholders...)
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
	query := "select p.id, title, content, name, user_id, height, width from pin_images join (select id, title, content, user_id, height, width " +
		"from (pins join pins_vectors_content on pins.id = pins_vectors_content.idv) " +
		"where \"vec\" @@ plainto_tsquery($1) union select pins.id, pins.title, pins.content, user_id, height, width from pins " +
		"where lower(title) like lower('%' || $1 || '%')) p " +
		"on p.id = pin_images.id "
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, title)
	if last > 0 {
		i++
		query += "where p.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by p.id desc limit 15"
	rows, err := r.db.Query(query, placeholders...)
	if err != nil {
		config.Lg("search", "GetPinsSearch").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	pins := make([]domain.Pin, 0, 15)
	for rows.Next() {
		pin := domain.Pin{}

		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId, &pin.PictureHeight, &pin.PictureWidth); err != nil {
			config.Lg("search", "GetPinSearch").Error(err.Error())
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}

func (r *Repository) GetBoardsByTitle(title string, last int) ([]domain.Board, error) {
	query := "select b.id, title, content, username from users join (" +
		"select id, title, content, user_id from (boards  join boards_vectors_content on boards.id = boards_vectors_content.idv) " +
		"where \"vec\" @@ plainto_tsquery($1) union select id, title, content, user_id from boards " +
		"where lower(title) like lower('%' || $1 || '%')) b on b.user_id = users.id "
	i := 1
	var placeholders []interface{}
	placeholders = append(placeholders, title)
	if last > 0 {
		i++
		query += " where b.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by b.id desc limit 15"
	rows, err := r.db.Query(query, placeholders...)
	if err != nil {
		config.Lg("search", "GetPinsSearch").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	boards := make([]domain.Board, 0, 15)
	for rows.Next() {
		board := domain.Board{}

		if err := rows.Scan(&board.ID, &board.Title, &board.Content, &board.Username); err != nil {
			config.Lg("search", "GetPinSearch").Error(err.Error())
			return nil, err
		}
		boards = append(boards, board)
	}
	return boards, nil
}
