package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
)

type Usecase struct {
	repo  notifications.IRepository
}

func NewUsecase(r notifications.IRepository) *Usecase {
	return &Usecase{
		repo:  r,
	}
}


