package feed

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type IRepository interface {
	GetFeed(userId int, last int) ([]domain.Pin, error)
	GetSubFeed(userId int, last int) ([]domain.Pin, error)
}
