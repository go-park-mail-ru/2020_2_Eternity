package utils

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"net/url"
	"path/filepath"
)

func GetUrlImg(imgName string) string {
	imgUrl := url.URL{
		Scheme: config.Conf.Web.Server.Protocol,
		Host:   config.Conf.Web.Server.Host,
		Path:   filepath.Join(config.Conf.Web.Static.UrlImg, imgName),
	}

	return imgUrl.String()
}
