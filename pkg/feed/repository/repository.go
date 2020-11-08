package repository

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type Repository struct {
	db database.IDbConn
}

func NewRepo(db database.IDbConn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetFeed(userId int) ([]domain.Pin, error) {
	rows, err := r.db.Query(context.Background(), "select pins.id, title, content, name, user_id "+
		"from pins join pin_images on pins.id = pin_images.pin_id order by pins.id desc")
	if err != nil {
		config.Lg("feed", "GetFeed").Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var pins []domain.Pin
	for rows.Next() {
		pin := domain.Pin{}
		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId); err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}
