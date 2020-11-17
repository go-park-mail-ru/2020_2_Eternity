package feed

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IUseCase interface {
	GetFeed(userId int, last int) ([]domain.PinResp, error)
}
