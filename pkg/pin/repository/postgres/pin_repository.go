package postgres

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}

func (r *Repository) StorePin(p *domain.Pin) error {
	err := r.dbConn.QueryRow(
		context.Background(),
		"with rows as "+
			"(insert into pins (title, content, user_id) "+
			"values($1, $2, $3) returning id) "+
			"insert into pin_images (name, pin_id) values($4, (select id from rows)) returning pin_id;",
		p.Title, p.Content, p.UserId, p.PictureName).Scan(&p.Id)

	if err != nil {
		config.Lg("pin", "pin.StorePin").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetPin(id int) (domain.Pin, error) {
	row := r.dbConn.QueryRow(
		context.Background(),
		"select pins.id, title, content, name, user_id "+
			"from pins join pin_images "+
			"on pins.id = pin_images.pin_id "+
			"where pins.id=$1;",
		id)

	p := domain.Pin{}
	if err := row.Scan(&p.Id, &p.Title, &p.Content, &p.PictureName, &p.UserId); err != nil {
		config.Lg("pin", "pin.GetPin").Error(err.Error())
		return domain.Pin{}, err
	}
	p.Id = id

	return p, nil
}

func (r *Repository) GetPinList(username string) ([]domain.Pin, error) {
	rows, err := r.dbConn.Query(
		context.Background(),
		"select pins.id, title, content, name, user_id "+
			"from pins join pin_images "+
			"on pins.id = pin_images.pin_id "+
			"where user_id in (select id from users where lower(username) = lower($1))",
		username)
	if err != nil {
		config.Lg("pin", "pin.GetPinList").Error(err.Error())
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

func (r *Repository) GetPinBoardList(boardId int) ([]domain.Pin, error) {
	rows, err := r.dbConn.Query(
		context.Background(),
		"select res.id, title, content, name, user_id from (pins join boards_pins on pins.id = boards_pins.pin_id)"+
			" res join pin_images on res.id = pin_images.pin_id"+
			" where res.board_id=$1;", boardId)
	if err != nil {
		config.Lg("pin", "pin.GetPinBoardList").Error(err.Error())
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
