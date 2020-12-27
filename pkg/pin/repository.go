package pin

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"mime/multipart"
)

type IRepository interface {
	StorePin(p *domain.Pin) error
	GetPin(id int) (domain.Pin, error)
	GetPinList(username string) ([]domain.Pin, error)

	GetPinBoardList(boardId int, limit int) ([]domain.Pin, error)
}

type IStorage interface {
	SaveUploadedFile(file *multipart.FileHeader, filename *string) error
	GetImageSizes(filename string) (int, int, error)
}
