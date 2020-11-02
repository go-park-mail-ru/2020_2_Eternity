package filestorage

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Storage struct {
	conf *config.Config
}

func NewStorage(c *config.Config) *Storage {
	return &Storage{
		conf: c,
	}
}

func (s *Storage) SaveUploadedFile(file *multipart.FileHeader, filename string) error {
	if err := os.MkdirAll(s.conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFile").
			Error("MkAllDir: ", err.Error())
		return err
	}

	src, err := file.Open()
	if err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFile").
			Error("MkAllDir: ", err.Error())
		return err
	}
	defer src.Close()

	dst := filepath.Join(s.conf.Web.Static.DirImg, filename)

	out, err := os.Create(dst)
	if err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFile").
			Error("MkAllDir: ", err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
