package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
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

func TestUsecase_GetUsersByName(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_search.NewMockIRepository(ctrl)
	usecase := NewUsecase(r)

	username := "21savage"
	last := 234
	r.EXPECT().GetUsersByName(username, last).Return([]domain.UserSearch{{Username: username}}, nil)
	users, err := usecase.GetUsersByName(username, last)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
}

func TestUsecase_GetPinsByTitle(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_search.NewMockIRepository(ctrl)
	usecase := NewUsecase(r)

	title := "21savage"
	last := 234
	r.EXPECT().GetPinsByTitle(title, last).Return([]domain.Pin{{Content: title}}, nil)
	pins, err := usecase.GetPinsByTitle(title, last)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pins))
}

func TestUsecase_GetPinsByTitleF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_search.NewMockIRepository(ctrl)
	usecase := NewUsecase(r)

	title := "21savage"
	last := 234
	r.EXPECT().GetPinsByTitle(title, last).Return(nil, errors.New(""))
	_, err := usecase.GetPinsByTitle(title, last)
	assert.Error(t, err)
}
