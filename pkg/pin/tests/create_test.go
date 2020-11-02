package pintests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"testing"
	"time"
)

var (
	pinCreate = pin.Pin{
		Title:       "t5",
		Content:     "c5",
		PictureName: "p5",
		UserId:      u.ID,
	}
)

func prepareCreatePin(t *testing.T) {
	err := config.Db.QueryRow(
		"insert into users(id, username, email, password, birthdate, reg_date, avatar) "+
			"values($1, $2, $3, $4, $5, $6, $7) returning id",
		u.ID, u.Username, u.Email, u.Password, u.BirthDate, time.Now(), "ava").Scan(&u.ID)

	log.Println("UserId : ", u.ID)

	if err != nil {
		assert.Fail(t, "Can't prepareGetPin DB: "+err.Error())
	}
}

func cleanupCreatePin(t *testing.T) {

	_, err := config.Db.Exec(
		"delete from pins "+
			"where user_id = $1",
		u.ID,
	)

	if err != nil {
		assert.Fail(t, "Can't cleanupGetPin DB: "+err.Error())
	}

	_, err = config.Db.Exec(
		"delete from users "+
			"where username = $1",
		u.Username,
	)

	if err != nil {
		assert.Fail(t, "Can't cleanupGetPin DB: "+err.Error())
	}
}

func createMultipartBody() (*bytes.Buffer, string, error) {
	buff := bytes.Buffer{}
	w := multipart.NewWriter(&buff)

	filePart, err := w.CreateFormFile("img", "filename")
	if err != nil {
		return nil, "", err
	}

	_, err = fmt.Fprintf(filePart, "file content")
	if err != nil {
		return nil, "", err
	}

	mH := textproto.MIMEHeader{}
	mH.Add("Content-Disposition", "form-data; name=\"data\";")
	mH.Add("Content-Type", "application/json")

	jsonPart, err := w.CreatePart(mH)
	if err != nil {
		return nil, "", err
	}

	js, err := json.Marshal(pinCreate)
	if err != nil {
		return nil, "", err
	}

	_, err = fmt.Fprintf(jsonPart, string(js))
	if err != nil {
		return nil, "", err
	}

	err = w.Close()
	if err != nil {
		return nil, "", err
	}

	return &buff, w.Boundary(), nil
}

func TestCreatePin(t *testing.T) {
	prepareCreatePin(t)
	defer cleanupCreatePin(t)

	body, boundary, err := createMultipartBody()
	if err != nil {
		assert.Fail(t, "Error while creating multipart")
	}

	resp, err := http.Post(fmt.Sprintf("%s/user/pin", ts.URL),
		"multipart/form-data; boundary="+boundary,
		body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	if err != nil || resp == nil {
		assert.Fail(t, "Error while request")
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail(t, "Error while Readall")
	}

	recvPinApi := api.PinResp{}
	if err := json.Unmarshal(recBody, &recvPinApi); err != nil {
		assert.Fail(t, "Error while unmarshal")
	}

	assert.Equal(t, recvPinApi.Title, pinCreate.Title)
	assert.Equal(t, recvPinApi.Content, pinCreate.Content)
	assert.Equal(t, recvPinApi.UserId, pinCreate.UserId)
}
