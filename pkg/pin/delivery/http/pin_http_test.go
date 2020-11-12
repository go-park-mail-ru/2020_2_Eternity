package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_database "github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_pin "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/mock"
	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strconv"
	"strings"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var (
	userId   = 3
	pinId    = 4
	username = "username123"
	boardId  = 6

	pinReq = domain.PinReq{
		Title:   "title",
		Content: "content",
	}

	pinResp = domain.PinResp{
		Id:      1,
		Title:   pinReq.Title,
		Content: pinReq.Content,
		ImgLink: "link",
		UserId:  userId,
	}

	pinRespMany = []domain.PinResp{
		{
			Id:      3,
			Title:   "tittle123",
			Content: "content14",
			ImgLink: "link4213",
			UserId:  userId,
		},
		{
			Id:      4,
			Title:   "tittle13",
			Content: "content4",
			ImgLink: "link13",
			UserId:  userId,
		},
	}
)

func createMultipartBody(pr domain.PinReq) (*bytes.Buffer, string, error) {
	buff := bytes.Buffer{}
	w := multipart.NewWriter(&buff)

	filePart, err := w.CreateFormFile("img", "filename")
	if err != nil {
		return nil, "", err
	}

	_, err = fmt.Fprintf(filePart, "file content")
	if err != nil {
		return nil, "", err
	}

	mH := textproto.MIMEHeader{}
	mH.Add("Content-Disposition", "form-data; name=\"data\";")
	mH.Add("Content-Type", "application/json")

	jsonPart, err := w.CreatePart(mH)
	if err != nil {
		return nil, "", err
	}

	js, err := json.Marshal(pr)
	if err != nil {
		return nil, "", err
	}

	_, err = fmt.Fprintf(jsonPart, string(js))
	if err != nil {
		return nil, "", err
	}

	err = w.Close()
	if err != nil {
		return nil, "", err
	}

	return &buff, w.Boundary(), nil
}

func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

func TestCreatePinSuccess(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	buff, boundary, err := createMultipartBody(pinReq)
	require.Nil(t, err)
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", buff)
	gCtx.Request.Header.Set("Content-Type", contentType)
	gCtx.Set("info", userId)

	mockUsecase := mock_pin.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		CreatePin(gomock.Eq(&pinReq), gomock.Any(), gomock.Eq(userId)).
		Return(pinResp, nil).
		Times(1)

	h.CreatePin(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	pResp := domain.PinResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &pResp))
	assert.Equal(t, pResp, pinResp)

	// Fail

	mockUsecase.EXPECT().
		CreatePin(gomock.Eq(&pinReq), gomock.Any(), gomock.Eq(userId)).
		Return(pinResp, errors.New("")).
		Times(1)

	h.CreatePin(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusInternalServerError)
}

func TestCreatePinFail(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", nil)

	mockUsecase := mock_pin.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	mockUsecase.EXPECT().
		CreatePin(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// No claims

	h.CreatePin(gCtx)
	assert.Equal(t, gCtx.Writer.Status(), http.StatusUnauthorized)

	// Bind

	gCtx.Set("info", userId)
	h.CreatePin(gCtx)
	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

	// Valid

	invalidPinReq := pinReq
	invalidPinReq.Content = strings.Repeat("s", 600)
	buff, boundary, err := createMultipartBody(invalidPinReq)
	require.Nil(t, err)
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)
	gCtx.Request = httptest.NewRequest("POST", "/", buff)
	gCtx.Request.Header.Set("Content-Type", contentType)

	h.CreatePin(gCtx)
	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

}

func TestGetPinById(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)
	gCtx.Params = append(gCtx.Params, gin.Param{Key: PinIdParam, Value: strconv.Itoa(pinId)})

	mockUsecase := mock_pin.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		GetPin(gomock.Eq(pinId)).
		Return(pinResp, nil).
		Times(1)

	h.GetPin(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	pResp := domain.PinResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &pResp))
	assert.Equal(t, pResp, pinResp)

	// Fail

	mockUsecase.EXPECT().
		GetPin(gomock.Eq(pinId)).
		Return(pinResp, errors.New("")).
		Times(1)

	h.GetPin(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusUnprocessableEntity)

	// no param
	gCtx.Params = gin.Params{}

	mockUsecase.EXPECT().
		GetPin(gomock.Eq(pinId)).
		Return(pinResp, errors.New("")).
		Times(0)

	h.GetPin(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)
}

func TestGetAllPins(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)
	gCtx.Params = append(gCtx.Params, gin.Param{Key: UsernameParam, Value: username})

	mockUsecase := mock_pin.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		GetPinList(gomock.Eq(username)).
		Return(pinRespMany, nil).
		Times(1)

	h.GetAllPins(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	psResp := []domain.PinResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &psResp))
	assert.Equal(t, psResp, pinRespMany)

	// Fail

	mockUsecase.EXPECT().
		GetPinList(gomock.Eq(username)).
		Return(pinRespMany, errors.New("")).
		Times(1)

	h.GetAllPins(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)
}

func TestGetPinsFromBoard(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)
	gCtx.Params = append(gCtx.Params, gin.Param{Key: BoardIdParam, Value: strconv.Itoa(boardId)})

	mockUsecase := mock_pin.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		GetPinBoardList(gomock.Eq(boardId)).
		Return(pinRespMany, nil).
		Times(1)

	h.GetPinsFromBoard(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	psResp := []domain.PinResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &psResp))
	assert.Equal(t, psResp, pinRespMany)

	// Fail

	mockUsecase.EXPECT().
		GetPinBoardList(gomock.Eq(boardId)).
		Return(pinRespMany, errors.New("")).
		Times(1)

	h.GetPinsFromBoard(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

	// no param

	gCtx.Params = gin.Params{}

	mockUsecase.EXPECT().
		GetPinBoardList(gomock.Eq(boardId)).
		Times(0)

	h.GetPinsFromBoard(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

}


func TestCreateRoutes(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	_, r := gin.CreateTestContext(writerResp)

	mockDatabase := mock_database.NewMockIDbConn(mockCtr)
	AddPinRoutes(r, mockDatabase, bluemonday.NewPolicy(), config.Conf)
}