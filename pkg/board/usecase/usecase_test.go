package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_board "github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestBoardUsecase_CreateBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	testBoard := api.CreateBoard{
		Title:   "my board",
		Content: "12345678 test",
	}
	b := &domain.Board{}
	userMockRepository.EXPECT().CreateBoard(1, &testBoard).Return(b, nil)
	if _, err := usecase.CreateBoard(1, &testBoard); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestBoardUsecase_GetBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	b := &domain.Board{}
	userMockRepository.EXPECT().GetBoard(1).Return(b, nil)
	if _, err := usecase.GetBoard(1); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestBoardUsecase_GetAllBoardsByUser(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "21savage"
	var b []domain.Board
	userMockRepository.EXPECT().GetAllBoardsByUser(username).Return(b, nil)
	if _, err := usecase.GetAllBoardsByUser(username); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestBoardUsecase_CheckOwner(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	userMockRepository.EXPECT().CheckOwner(1, 1).Return(nil)
	if err := usecase.CheckOwner(1, 1); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestBoardUsecase_AttachPin(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	userMockRepository.EXPECT().AttachPin(1, 1).Return(nil)
	if err := usecase.AttachPin(1, 1); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestBoardUsecase_DetachPin(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_board.NewMockIUsecase(ctrl)
	usecase := NewUsecase(userMockRepository)

	userMockRepository.EXPECT().DetachPin(1, 1).Return(nil)
	if err := usecase.DetachPin(1, 1); err != nil {
		assert.Equal(t, err, nil)
	}
}
