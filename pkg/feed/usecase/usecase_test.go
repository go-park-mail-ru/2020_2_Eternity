package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_feed "github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestHandler_Feed(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_feed.NewMockIRepository(ctrl)
	uc := NewUseCase(r)

	r.EXPECT().GetFeed(0, 234).Return([]domain.Pin{{Title: "12345"}}, nil)

	pins, err := uc.GetFeed(0, 234)
	assert.NoError(t, err)
	assert.Equal(t, len(pins), 1)
}

func TestHandler_FeedF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_feed.NewMockIRepository(ctrl)
	uc := NewUseCase(r)

	r.EXPECT().GetFeed(0, 234).Return(nil, errors.New(""))

	pins, err := uc.GetFeed(0, 234)
	assert.Error(t, err)
	assert.Equal(t, []domain.PinResp([]domain.PinResp(nil)), pins)
}
