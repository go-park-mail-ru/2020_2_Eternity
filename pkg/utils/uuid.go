package utils

import (
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


//func GenerateUrlAvatar(filename string) string {
//	return "http://" + config.Conf.Web.Server.Host + "/images/avatar/" + filename
//}

func GenerateUrlAvatar(filename string) string {
	return "/api/images/avatar/" + filename
}