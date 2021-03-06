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

func NewRepo(db database.IDbConn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetFeed(userId int, last int) ([]domain.Pin, error) {
	query := "select pins.id, title, content, name, user_id, height, width " +
		"from pins join pin_images on pins.id = pin_images.pin_id"
	var placeholders []interface{}
	i := 0
	if last > 0 {
		i++
		query += " where pins.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by pins.id desc limit 15"
	rows, err := r.db.Query(query, placeholders...)
	if err != nil {
		config.Lg("feed", "GetFeed").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	pins := make([]domain.Pin, 0, 15)
	for rows.Next() {
		pin := domain.Pin{}
		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId, &pin.PictureHeight, &pin.PictureWidth); err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}

func (r *Repository) GetSubFeed(userId int, last int) ([]domain.Pin, error) {
	query := "select p.id, title, content, name, user_id, height, width " +
		"from (pins join follows on user_id = follows.id2 and follows.id1 = $1) p " +
		"join pin_images on p.id = pin_images.pin_id "
	var placeholders []interface{}
	placeholders = append(placeholders, userId)
	i := 1
	if last > 0 {
		i++
		query += " where p.id < $" + strconv.Itoa(i)
		placeholders = append(placeholders, last)
	}
	query += " order by p.id desc limit 15"
	rows, err := r.db.Query(query, placeholders...)
	if err != nil {
		config.Lg("feed", "GetSubFeed").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	pins := make([]domain.Pin, 0, 15)
	for rows.Next() {
		pin := domain.Pin{}
		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId, &pin.PictureHeight, &pin.PictureWidth); err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}
