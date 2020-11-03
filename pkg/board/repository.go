package board

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type IRepository interface {
	CreateBoard(userId int, b *api.CreateBoard) (*domain.Board, error)
	GetBoard(id int) (*domain.Board, error)
	GetAllBoardsByUser(username string) ([]domain.Board, error)
}
