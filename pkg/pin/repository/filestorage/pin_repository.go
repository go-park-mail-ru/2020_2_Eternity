package filestorage

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	conf *config.Config
}

func NewStorage(c *config.Config) *Storage {
	return &Storage{
		conf: c,
	}
}

func (s *Storage) GetImageSizes(filename string) (int, int, error) {
	src, err := os.Open(filepath.Join(s.conf.Web.Static.DirImg, filename))
	if err != nil {
		config.Lg("pin", "GetSize").Error(err.Error())
		return 0, 0, err
	}
	defer src.Close()

	img, _, err := image.DecodeConfig(src)
	if err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFileImage").Error(err.Error())
		return 0, 0, err
	}

	log.Println(img.Height, img.Width)
	return img.Height, img.Width, nil
}

func (s *Storage) SaveUploadedFile(file *multipart.FileHeader, filename *string) error {
	if err := os.MkdirAll(s.conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFile").
			Error("MkAllDir: ", err.Error())
		return err
	}

	lastIdx := strings.LastIndex(file.Filename, ".")
	if lastIdx == -1 {
		return errors.New("need format file")
	}

	format := []byte(file.Filename)[lastIdx:]

	src, err := file.Open()
	if err != nil {
		config.Lg("pin_filestorage", "SaveUploadedFile").
			Error("MkAllDir: ", err.Error())
		return err
	}
	defer src.Close()

	*filename = *filename + string(format)

	dst := filepath.Join(s.conf.Web.Static.DirImg, *filename)

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
