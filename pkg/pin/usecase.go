package pin

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"mime/multipart"
)

type IUsecase interface {
	CreatePin(pin *domain.PinReq, file *multipart.FileHeader, userId int) (domain.PinResp, error)
	GetPinList(username string) ([]domain.PinResp, error)

	GetPinBoardList(boardId int) ([]domain.PinResp, error)
}
