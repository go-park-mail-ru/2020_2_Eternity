package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_feed "github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

	uc := mock_feed.NewMockIUseCase(ctrl)
	handler := NewHandler(uc)

	path := "/feed"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?last=234", nil)

	c, r := gin.CreateTestContext(w)
	uc.EXPECT().GetFeed(0, 234).Return([]domain.PinResp{}, nil)
	r.GET(path, handler.GetFeed)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_FeedF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_feed.NewMockIUseCase(ctrl)
	handler := NewHandler(uc)

	path := "/feed"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?last=234", nil)

	c, r := gin.CreateTestContext(w)
	uc.EXPECT().GetFeed(0, 234).Return([]domain.PinResp{}, errors.New(""))
	r.GET(path, handler.GetFeed)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 500, c.Writer.Status())
}
