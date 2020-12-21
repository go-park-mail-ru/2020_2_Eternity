package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	sc "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	mock_search "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search/mock"
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

func TestHandler_SearchUser(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=user&content=21s&last=234", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetUsersByName(gomock.Any(), &sc.UserSearch{
		Username: "21s",
		Last:     int32(234),
	}).Return(&sc.Users{
		Users: []*sc.User{
			{Username: "21savage"},
		},
	}, nil)
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchUserF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=user&content=22s&last=234", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetUsersByName(gomock.Any(), &sc.UserSearch{
		Username: "22s",
		Last:     int32(234),
	}).Return(nil, errors.New(""))
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchPin(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=pin&content=album&last=234", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetPinsByTitle(gomock.Any(), &sc.PinSearch{
		Title: "album",
		Last:  int32(234),
	}).Return(&sc.Pins{
		Pins: []*sc.Pin{
			{Title: "album drop"},
		},
	}, nil)
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchPinF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=pin&content=al", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetPinsByTitle(gomock.Any(), &sc.PinSearch{
		Title: "al",
		Last:  0,
	}).Return(nil, errors.New(""))
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=board&content=album&last=234", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetBoardsByTitle(gomock.Any(), &sc.BoardSearch{
		Title: "album",
		Last:  int32(234),
	}).Return(&sc.Boards{
		Boards: []*sc.Board{
			{Title: "album drop"},
		},
	}, nil)
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchBoardF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=board&content=al", nil)

	c, r := gin.CreateTestContext(w)
	ss.EXPECT().GetBoardsByTitle(gomock.Any(), &sc.BoardSearch{
		Title: "al",
		Last:  0,
	}).Return(nil, errors.New(""))
	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_SearchF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ss := mock_search.NewMockSearchServiceClient(ctrl)
	handler := NewHandler(ss)

	path := "/search"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", path+"?type=pin", nil)

	c, r := gin.CreateTestContext(w)

	r.GET(path, handler.Search)

	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}
