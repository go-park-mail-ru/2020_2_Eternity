package search

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IRepository interface {
	GetUsersByName(username string, last int) ([]domain.UserSearch, error)
	GetPinsByTitle(title string, last int) ([]domain.Pin, error)
}
