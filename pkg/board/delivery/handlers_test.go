package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_board "github.com/go-park-mail-ru/2020_2_Eternity/pkg/board/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var p = bluemonday.UGCPolicy()

func mid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("info", 1)
		c.Next()
	}
}

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewTestConfig()
	return true
}()

func TestHandler_CreateBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board"

	b := api.CreateBoard{
		Title:   "Album",
		Content: "Dropped hey",
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	MockUsecase.EXPECT().CreateBoard(1, &b).Return(&domain.Board{}, nil)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.CreateBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_CreateBoardU(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board"

	req, _ := http.NewRequest("POST", path, nil)

	c, r := gin.CreateTestContext(w)

	r.POST(path, Handler.CreateBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 401, c.Writer.Status())
}

func TestHandler_CreateBoardF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board"

	b := api.CreateBoard{
		Title:   "A",
		Content: "D",
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	MockUsecase.EXPECT().CreateBoard(1, &b).Return(&domain.Board{}, errors.New(""))

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.CreateBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 422, c.Writer.Status())
}

func TestHandler_GetBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board/12"

	MockUsecase.EXPECT().GetBoard(12).Return(&domain.Board{}, nil)

	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/board/:id", Handler.GetBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_GetBoardF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board/adas"

	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/board/:id", Handler.GetBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 400, c.Writer.Status())
}

func TestHandler_GetBoardW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/board/12"

	MockUsecase.EXPECT().GetBoard(12).Return(&domain.Board{}, errors.New(""))

	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/board/:id", Handler.GetBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 422, c.Writer.Status())
}

func TestHandler_GetBoards(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/boards/21savage"

	var bs []domain.Board

	MockUsecase.EXPECT().GetAllBoardsByUser("21savage").Return(bs, nil)

	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/boards/:username", Handler.GetAllBoardsbyUser)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_GetBoardsW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/boards/21s"

	var bs []domain.Board

	MockUsecase.EXPECT().GetAllBoardsByUser("21s").Return(bs, errors.New(""))

	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/boards/:username", Handler.GetAllBoardsbyUser)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 422, c.Writer.Status())
}

func TestHandler_AttachPinToBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/attach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 1,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	MockUsecase.EXPECT().CheckOwner(1, b.BoardID).Return(nil)
	MockUsecase.EXPECT().AttachPin(b.BoardID, b.PinID).Return(nil)
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.AttachPinToBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_AttachPinToBoardF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/attach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 2,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	MockUsecase.EXPECT().CheckOwner(1, b.BoardID).Return(errors.New(""))
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.AttachPinToBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 400, c.Writer.Status())
}

func TestHandler_AttachPinToBoardW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/attach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 2,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	MockUsecase.EXPECT().CheckOwner(1, b.BoardID).Return(nil)
	MockUsecase.EXPECT().AttachPin(b.BoardID, b.PinID).Return(errors.New(""))
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.AttachPinToBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 400, c.Writer.Status())
}

func TestHandler_DetachPinFromBoard(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/detach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 1,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	MockUsecase.EXPECT().CheckOwner(1, b.BoardID).Return(nil)
	MockUsecase.EXPECT().DetachPin(b.BoardID, b.PinID).Return(nil)
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.DetachPinFromBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_DetachPinFromBoardU(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/detach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 2,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))
	c, r := gin.CreateTestContext(w)
	r.POST(path, Handler.DetachPinFromBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 401, c.Writer.Status())
}

func TestHandler_DetachPinFromBoardW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockUsecase := mock_board.NewMockIUsecase(ctrl)
	Handler := NewHandler(MockUsecase, p)

	w := httptest.NewRecorder()

	path := "/detach"

	b := api.AttachDetachPin{
		PinID:   1,
		BoardID: 2,
	}

	body, err := json.Marshal(b)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	MockUsecase.EXPECT().CheckOwner(1, b.BoardID).Return(nil)
	MockUsecase.EXPECT().DetachPin(b.BoardID, b.PinID).Return(errors.New(""))
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, Handler.DetachPinFromBoard)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 400, c.Writer.Status())
}
