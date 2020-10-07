package utils

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/google/uuid"
	"strings"
)

func RandomUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.Replace(u.String(), "-", "", -1), err
}

func GenerateUrlAvatar(filename string) string {
	return "http://" + config.Conf.Web.Server.Address + ":" + config.Conf.Web.Server.Port + "/images/avatar/" + filename
}
