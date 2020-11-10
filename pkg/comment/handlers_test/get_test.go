package handlers_test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	mock_comment "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var (
	commentGetAllCase = []comment.Comment{
		{
			Id:      1,
			Path:    []int32{1},
			Content: "content",
			PinId:   1,
			UserId:  1,
		},
		{
			Id:      2,
			Path:    []int32{1, 2},
			Content: "content2",
			PinId:   2,
			UserId:  2,
		},
	}

	commentGetIdCase = commentGetAllCase[0]
)

func TestGetAllComments(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: comment.PinIdParam, Value: "4"})

	ctx.Set()
	ctx.Param()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoComment := mock_comment.NewMockRepoComment(ctrl)
	responder := NewMockResponder(mockRepoComment)

	// Success

	mockRepoComment.
		EXPECT().
		GetAllComments(gomock.Eq(4)).
		Times(1).
		Return(commentGetAllCase, nil)

	responder.GetComments(ctx)

	assert.Equal(t, 200, ctx.Writer.Status())

	commentsApi := []api.CommentResp{}
	assert.Nil(t, json.Unmarshal(writer.Body.Bytes(), &commentsApi))
	assert.Equal(
		t,
		[]api.CommentResp{
			{
				Id:      commentGetAllCase[0].Id,
				Path:    commentGetAllCase[0].Path,
				Content: commentGetAllCase[0].Content,
				PinId:   commentGetAllCase[0].PinId,
				UserId:  commentGetAllCase[0].UserId,
			},
			{
				Id:      commentGetAllCase[1].Id,
				Path:    commentGetAllCase[1].Path,
				Content: commentGetAllCase[1].Content,
				PinId:   commentGetAllCase[1].PinId,
				UserId:  commentGetAllCase[1].UserId,
			},
		},
		commentsApi)

	// Fail - repo returns error

	mockRepoComment.
		EXPECT().
		GetAllComments(gomock.Eq(4)).
		Times(1).
		Return([]comment.Comment{}, errors.New(""))

	responder.GetComments(ctx)
	assert.Equal(t, 400, ctx.Writer.Status())

	// Fail - path params not given

	ctx.Params = gin.Params{}
	mockRepoComment.
		EXPECT().
		GetAllComments(gomock.Eq(4)).
		Times(0)

	responder.GetComments(ctx)
	assert.Equal(t, 400, ctx.Writer.Status())
}

func TestGetCommentById(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: comment.CommentIdParam, Value: "4"})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoComment := mock_comment.NewMockRepoComment(ctrl)
	responder := NewMockResponder(mockRepoComment)

	// Success

	mockRepoComment.
		EXPECT().
		GetComment(gomock.Eq(4)).
		Times(1).
		Return(commentGetIdCase, nil)

	responder.GetCommentById(ctx)

	assert.Equal(t, 200, ctx.Writer.Status())

	commentApi := api.CommentResp{}
	assert.Nil(t, json.Unmarshal(writer.Body.Bytes(), &commentApi))
	assert.Equal(
		t,
		api.CommentResp{
			Id:      commentGetAllCase[0].Id,
			Path:    commentGetAllCase[0].Path,
			Content: commentGetAllCase[0].Content,
			PinId:   commentGetAllCase[0].PinId,
			UserId:  commentGetAllCase[0].UserId,
		},
		commentApi)

	// Fail - repo returned error

	mockRepoComment.
		EXPECT().
		GetComment(gomock.Eq(4)).
		Times(1).
		Return(comment.Comment{}, errors.New(""))

	responder.GetCommentById(ctx)
	assert.Equal(t, 400, ctx.Writer.Status())

	// Fail - path params not given

	ctx.Params = gin.Params{}
	mockRepoComment.
		EXPECT().
		GetComment(gomock.Eq(4)).
		Times(0)

	responder.GetCommentById(ctx)
	assert.Equal(t, 400, ctx.Writer.Status())
}
