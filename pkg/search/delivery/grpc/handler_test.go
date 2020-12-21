package grpc

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	sc "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	mock_search "github.com/go-park-mail-ru/2020_2_Eternity/pkg/search/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestSearchHandler_GetUsersByName(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	us := &sc.UserSearch{Username: "21savage", Last: int32(234)}

	uc.EXPECT().GetUsersByName(us.Username, int(us.Last)).Return([]domain.UserSearch{{Username: "21savage"}}, nil)

	users, err := handler.GetUsersByName(context.Background(), us)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users.Users))

}

func TestSearchHandler_GetUsersByNameF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	us := &sc.UserSearch{Username: "21savage", Last: int32(234)}

	uc.EXPECT().GetUsersByName(us.Username, int(us.Last)).Return(nil, errors.New(""))

	_, err := handler.GetUsersByName(context.Background(), us)
	assert.Error(t, err)

}

func TestSearchHandler_GetPinsByTitle(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	ps := &sc.PinSearch{Title: "album", Last: int32(234)}

	uc.EXPECT().GetPinsByTitle(ps.Title, int(ps.Last)).Return([]domain.PinResp{{Title: "album drop"}}, nil)

	pins, err := handler.GetPinsByTitle(context.Background(), ps)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pins.Pins))

}

func TestSearchHandler_GetPinsByTitleF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	ps := &sc.PinSearch{Title: "a", Last: int32(234)}

	uc.EXPECT().GetPinsByTitle(ps.Title, int(ps.Last)).Return(nil, errors.New(""))

	_, err := handler.GetPinsByTitle(context.Background(), ps)
	assert.Error(t, err)
}

func TestSearchHandler_GetBoardsByTitle(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	bs := &sc.BoardSearch{
		Title: "test",
		Last:  int32(234),
	}
	uc.EXPECT().GetBoardsByTitle(bs.Title, int(bs.Last)).Return([]domain.Board{{Title: "test"}}, nil)

	boards, err := handler.GetBoardsByTitle(context.Background(), bs)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(boards.GetBoards()))
}

func TestSearchHandler_GetBoardsByTitleF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_search.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	bs := &sc.BoardSearch{
		Title: "test asdasd",
		Last:  int32(234),
	}
	uc.EXPECT().GetBoardsByTitle(bs.Title, int(bs.Last)).Return(nil, errors.New(""))

	_, err := handler.GetBoardsByTitle(context.Background(), bs)
	assert.Error(t, err)
}
