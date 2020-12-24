package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/board"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type BoardUsecase struct {
	repo board.IRepository
}

func NewUsecase(repo board.IRepository) *BoardUsecase {
	return &BoardUsecase{
		repo: repo,
	}
}

func (uc *BoardUsecase) CreateBoard(userId int, b *api.CreateBoard) (*domain.Board, error) {
	return uc.repo.CreateBoard(userId, b)
}

func (uc *BoardUsecase) GetBoard(id int) (*domain.Board, error) {
	return uc.repo.GetBoard(id)
}

func (uc *BoardUsecase) GetAllBoardsByUser(username string) ([]domain.Board, error) {
	return uc.repo.GetAllBoardsByUser(username)
}

func (uc *BoardUsecase) CheckOwner(userId int, boardId int) error {
	return uc.repo.CheckOwner(userId, boardId)
}

func (uc *BoardUsecase) AttachPin(boardId int, pinId int) error {
	return uc.repo.AttachPin(boardId, pinId)
}
func (uc *BoardUsecase) DetachPin(boardId int, pinId int) error {
	return uc.repo.DetachPin(boardId, pinId)
}

func (uc *BoardUsecase) GetBoardsPinNotAttach(userId, pinId int) ([]domain.Board, error) {
	return uc.repo.GetBoardsPinNotAttach(userId, pinId)
}
