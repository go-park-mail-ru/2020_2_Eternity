package search

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IUsecase interface {
	GetUsersByName(username string, last int) ([]domain.UserSearch, error)
	GetPinsByTitle(title string, last int) ([]domain.PinResp, error)
}
