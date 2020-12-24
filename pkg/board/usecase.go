package board

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type IUsecase interface {
	CreateBoard(userId int, b *api.CreateBoard) (*domain.Board, error)
	GetBoard(id int) (*domain.Board, error)
	GetAllBoardsByUser(username string) ([]domain.Board, error)

	CheckOwner(userId int, boardId int) error

	AttachPin(boardId int, pinId int) error
	DetachPin(boardId int, pinId int) error
	GetBoardsPinNotAttach(userId, pinId int) ([]domain.Board, error)
}
