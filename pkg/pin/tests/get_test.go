package pintests

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
	"time"
)

var (
	pinsGet = []pin.Pin{
		{
			Title:       "t",
			Content:     "c",
			PictureName: "p",
			UserId:      u.ID,
		},
		{
			Title:       "t2",
			Content:     "c2",
			PictureName: "p2",
			UserId:      u.ID,
		},
		{
			Title:       "t3",
			Content:     "c3",
			PictureName: "p3",
			UserId:      u.ID,
		},
	}
)

func prepareGetPin(t *testing.T) {
	err := config.Db.QueryRow(
		"insert into users(id, username, email, password, birthdate, reg_date, avatar) "+
			"values($1, $2, $3, $4, $5, $6, $7) returning id",
		u.ID, u.Username, u.Email, u.Password, u.BirthDate, time.Now(), "ava").Scan(&u.ID)

	log.Println("UserId : ", u.ID)

	if err != nil {
		assert.Fail(t, "Can't prepareGetPin DB: "+err.Error())
	}

	for i, p := range pinsGet {

		err = config.Db.QueryRow(
			"with rows as "+
				"(insert into pins (title, content, user_id) "+
				"values($1, $2, $3) returning id) "+
				"insert into pin_images (name, pin_id) values($4, (select id from rows)) returning pin_id;",
			p.Title, p.Content, p.UserId, p.PictureName).Scan(&pinsGet[i].Id)

		if err != nil {
			assert.Fail(t, "Can't prepareGetPin DB: "+err.Error())
		}
	}
}

func cleanupGetPin(t *testing.T) {
	for _, p := range pinsGet {
		_, err := config.Db.Exec(
			"delete from pins "+
				"where id = $1",
			p.Id,
		)

		if err != nil {
			assert.Fail(t, "Can't cleanupGetPin DB: "+err.Error())
		}
	}

	_, err := config.Db.Exec(
		"delete from users "+
			"where username = $1",
		u.Username,
	)

	if err != nil {
		assert.Fail(t, "Can't cleanupGetPin DB: "+err.Error())
	}
}

func createPinsApi() []api.PinResp {
	pinsApi := []api.PinResp{}

	for _, p := range pinsGet {
		imgUrl := url.URL{
			Scheme: config.Conf.Web.Server.Protocol,
			Host:   config.Conf.Web.Server.Host,
			Path:   filepath.Join(config.Conf.Web.Static.UrlImg, p.PictureName),
		}

		pinsApi = append(pinsApi, api.PinResp{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
			UserId:  p.UserId,
			ImgLink: imgUrl.String(),
		})
	}
	return pinsApi
}

func TestGetPin(t *testing.T) {
	prepareGetPin(t)
	defer cleanupGetPin(t)

	resp, err := http.Get(fmt.Sprintf("%s/user/pin", ts.URL))
	if err != nil || resp == nil {
		assert.Fail(t, "Error while request")
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, resp.Body)

	bodyBuff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail(t, "Error while readAll")
	}

	bodyExp, err := json.Marshal(createPinsApi())
	if err != nil {
		assert.Fail(t, "Error while readAll")
	}

	assert.Equal(t, string(bodyExp), string(bodyBuff))
}
