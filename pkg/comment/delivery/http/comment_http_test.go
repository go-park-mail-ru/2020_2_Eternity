package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_comment "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"

	"testing"
)

var (
//gCtx *gin.Context
//writerResp *httptest.ResponseRecorder
)

var (
	userId = 3

	commentReqRoot = domain.CommentCreateReq{
		IsRoot:   true,
		ParentId: 1,
		Content:  "content",
		PinId:    2,
	}

	commentRespRoot = domain.CommentResp{
		Id:      1,
		Path:    []int32{1},
		Content: commentReqRoot.Content,
		PinId:   commentReqRoot.PinId,
		UserId:  userId,
	}

	pinId     = 5
	commentId = 10

	commentRespOne = domain.CommentResp{
		Id:      2,
		Path:    []int32{1, 2},
		Content: "content",
		PinId:   12,
		UserId:  userId,
	}

	commentRespMany = []domain.CommentResp{
		{
			Id:      12,
			Path:    []int32{2},
			Content: "content2",
			PinId:   7,
			UserId:  userId,
		},
		{
			Id:      3,
			Path:    []int32{1, 3, 2},
			Content: "content3",
			PinId:   8,
			UserId:  userId,
		},
	}
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateCommentSuccess(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	buff, err := json.Marshal(commentReqRoot)
	require.Nil(t, err)
	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(buff))
	gCtx.Set("info", userId)

	mockUsecase := mock_comment.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		CreateComment(gomock.Eq(&commentReqRoot), gomock.Eq(userId)).
		Return(commentRespRoot, nil).
		Times(1)

	h.CreateComment(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	commResp := domain.CommentResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &commResp))
	assert.Equal(t, commResp, commentRespRoot)

	// Fail

	gCtx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(buff))

	mockUsecase.EXPECT().
		CreateComment(gomock.Eq(&commentReqRoot), gomock.Eq(userId)).
		Return(commentRespRoot, errors.New("")).
		Times(1)

	h.CreateComment(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)
}

func TestCreateCommentFail(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	buff, err := json.Marshal(commentReqRoot)
	require.Nil(t, err)
	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(buff))

	mockUsecase := mock_comment.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	mockUsecase.EXPECT().
		CreateComment(gomock.Any(), gomock.Any()).
		Times(0)

	// No claims

	h.CreateComment(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusUnauthorized)

	// bind

	gCtx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte("{")))
	gCtx.Set("info", userId)

	h.CreateComment(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

	// Validate

	wrongReq := commentRespRoot
	wrongReq.Content = strings.Repeat("s", 500)
	buff, err = json.Marshal(wrongReq)
	require.Nil(t, err)

	gCtx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(buff))

	h.CreateComment(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)
}

func TestGetPinComments(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", nil)
	gCtx.Params = append(gCtx.Params, gin.Param{Key: PinIdParam, Value: strconv.Itoa(pinId)})

	mockUsecase := mock_comment.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		GetPinComments(gomock.Eq(pinId)).
		Return(commentRespMany, nil).
		Times(1)

	h.GetPinComments(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	commsResp := []domain.CommentResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &commsResp))
	assert.Equal(t, commsResp, commentRespMany)

	// Fail

	mockUsecase.EXPECT().
		GetPinComments(gomock.Eq(pinId)).
		Return(commentRespMany, errors.New("")).
		Times(1)

	h.GetPinComments(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

}

func TestGetPinCommentsFail(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("POST", "/", nil)

	mockUsecase := mock_comment.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// no param

	mockUsecase.EXPECT().
		GetPinComments(gomock.Any()).
		Times(0)

	h.GetPinComments(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

}

func TestGetCommentById(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)
	gCtx.Params = append(gCtx.Params, gin.Param{Key: CommentIdParam, Value: strconv.Itoa(commentId)})

	mockUsecase := mock_comment.NewMockIUsecase(mockCtr)
	h := NewHandler(mockUsecase, bluemonday.NewPolicy())

	// Success

	mockUsecase.EXPECT().
		GetCommentById(gomock.Eq(commentId)).
		Return(commentRespOne, nil).
		Times(1)

	h.GetCommentById(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	commResp := domain.CommentResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &commResp))
	assert.Equal(t, commResp, commentRespOne)

	// Fail

	mockUsecase.EXPECT().
		GetCommentById(gomock.Eq(commentId)).
		Return(commentRespOne, errors.New("")).
		Times(1)

	h.GetCommentById(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusBadRequest)

}
