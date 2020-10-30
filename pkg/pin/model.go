package pin

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"log"
	"path/filepath"
)

type Pin struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PictureName string `json:"picture_name"`
	UserId      int    `json:"user_id"`
}

func (p *Pin) CreatePin() error {
	err := config.Db.QueryRow(
		"with rows as "+
			"(insert into pins (title, content, user_id) "+
			"values($1, $2, $3) returning id) "+
			"insert into pin_images (name, pin_id) values($4, (select id from rows)) returning pin_id;",
		p.Title, p.Content, p.UserId, p.PictureName).Scan(&p.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *Pin) GetPin(id int) error {
	row := config.Db.QueryRow(
		"select id, title, content, name, user_id "+
			"from pins join pin_images "+
			"on pins.id = pin_images.pin_id "+
			"where pins.id=$1;",
		id)

	if err := row.Scan(&p.Id, &p.Title, &p.Content, &p.PictureName, &p.UserId); err != nil {
		log.Print(err)
		return err
	}

	p.Id = id
	return nil
}

func GetPinList(userId int) ([]api.GetPin, error) {
	rows, err := config.Db.Query(
		"select pins.id, title, content, name, user_id "+
			"from pins join pin_images "+
			"on pins.id = pin_images.pin_id "+
			"where user_id=$1;",
		userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	pins := []api.GetPin{}
	for rows.Next() {
		pin := Pin{}
		if err := rows.Scan(&pin.Id, &pin.Title, &pin.Content, &pin.PictureName, &pin.UserId); err != nil {
			return nil, err
		}
		pins = append(pins, api.GetPin{
			Id:      pin.Id,
			Title:   pin.Title,
			Content: pin.Content,
			ImgLink: filepath.Join(config.Conf.Web.Static.UrlImg, pin.PictureName), // TODO full path
			UserId:  pin.UserId,
		})
	}

	return pins, nil
}
