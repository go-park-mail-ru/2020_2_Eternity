package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_notifications "github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications/usecase"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	userId = 32

	noteRespMany = []domain.NoteResp {
		{
			Id: 12,
			Type: usecase.NoteComment,
			EncodedData: []byte("123123"),
			IsRead: false,
		},
		{
			Id: 13,
			Type: usecase.NotePin,
			EncodedData: []byte("1223"),
			IsRead: false,
		},
	}

)

func TestGetUserNotes(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()


	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)
	gCtx.Set("info", userId)

	mockUsecase := mock_notifications.NewMockIUsecase(mockCtr)
	h := NewDelivery(mockUsecase)


	// Success

	mockUsecase.EXPECT().
		GetUserNotes(gomock.Eq(userId)).
		Return(noteRespMany, nil).
		Times(1)

	h.GetUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)
	nResp := []domain.NoteResp{}
	require.Nil(t, json.Unmarshal(writerResp.Body.Bytes(), &nResp))
	assert.Equal(t, noteRespMany, nResp)

	// Fail

	mockUsecase.EXPECT().
		GetUserNotes(gomock.Eq(userId)).
		Return(noteRespMany, errors.New("")).
		Times(1)

	h.GetUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusInternalServerError)

	// no claims


	gCtx, _ = gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)

	mockUsecase.EXPECT().
		GetUserNotes(gomock.Eq(userId)).
		Times(0)

	h.GetUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusUnauthorized)
}


func TestUpdUserNotes(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()


	writerResp := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("PUT", "/", nil)
	gCtx.Set("info", userId)

	mockUsecase := mock_notifications.NewMockIUsecase(mockCtr)
	h := NewDelivery(mockUsecase)


	// Success

	mockUsecase.EXPECT().
		UpdateUserNotes(gomock.Eq(userId)).
		Return(nil).
		Times(1)

	h.UpdateUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusOK)

	// Fail

	mockUsecase.EXPECT().
		UpdateUserNotes(gomock.Eq(userId)).
		Return(errors.New("")).
		Times(1)

	h.UpdateUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusInternalServerError)

	// no claims


	gCtx, _ = gin.CreateTestContext(writerResp)
	gCtx.Request = httptest.NewRequest("GET", "/", nil)

	mockUsecase.EXPECT().
		UpdateUserNotes(gomock.Eq(userId)).
		Times(0)

	h.UpdateUserNotes(gCtx)

	assert.Equal(t, gCtx.Writer.Status(), http.StatusUnauthorized)
}


//func TestCreateRoutes(t *testing.T) {
//	mockCtr := gomock.NewController(t)
//	defer mockCtr.Finish()
//
//	writerResp := httptest.NewRecorder()
//	_, r := gin.CreateTestContext(writerResp)
//
//	mockDatabase := mock_database.NewMockIDbConn(mockCtr)
//	AddNoteRoutes(r, mockDatabase)
//}
//
//
//func TestCreateNoteUsecase(t *testing.T) {
//	mockCtr := gomock.NewController(t)
//	defer mockCtr.Finish()
//
//
//	mockDatabase := mock_database.NewMockIDbConn(mockCtr)
//	CreateNoteUsecase(mockDatabase)
//}