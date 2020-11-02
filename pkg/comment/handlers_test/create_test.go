package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	mock_comment "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	commentCreateCase = comment.Comment{
		//Id: 0,
		//Path: []int32{},
		Content: "content",
		PinId:   1,
		UserId:  1,
	}

	commentApiCreateRootCase = api.CommentCreateReq{
		IsRoot:   true,
		ParentId: 1,
		Content:  commentCreateCase.Content,
		PinId:    commentCreateCase.PinId,
	}

	commentApiGetRootCase = api.CommentResp{
		Id:      3,
		Path:    []int32{1, 2, 3},
		Content: commentCreateCase.Content,
		PinId:   commentCreateCase.PinId,
		UserId:  commentCreateCase.UserId,
	}

	commentApiCreateChildCase = api.CommentCreateReq{
		IsRoot:   false,
		ParentId: 2,
		Content:  commentCreateCase.Content,
		PinId:    commentCreateCase.PinId,
	}

	commentApiGetChildCase = api.CommentResp{
		Id:      5,
		Path:    []int32{1, 3},
		Content: commentCreateCase.Content,
		PinId:   commentCreateCase.PinId,
		UserId:  commentCreateCase.UserId,
	}
)

type CreateContext struct {
	responder       *comment.ResponderComment
	mockRepoComment *mock_comment.MockRepoComment
	ctx             *gin.Context
	writer          *httptest.ResponseRecorder
	ctrl            *gomock.Controller
}

func createRequest(t *testing.T, bodyI interface{}) *http.Request {
	switch body := bodyI.(type) {
	case []byte:
		return httptest.NewRequest("GET", "/", bytes.NewBuffer(body))

	default:
		buff, err := json.Marshal(bodyI)
		require.Nil(t, err)
		return httptest.NewRequest("GET", "/", bytes.NewBuffer(buff))
	}
}

func prepareTestCreate(t *testing.T, reqBody interface{}) CreateContext {
	cc := CreateContext{}

	cc.writer = httptest.NewRecorder()

	cc.ctx, _ = gin.CreateTestContext(cc.writer)
	cc.ctx.Request = createRequest(t, reqBody)

	claims := jwthelper.Claims{
		Id: commentCreateCase.UserId,
	}
	cc.ctx.Set("info", claims.Id)

	cc.ctrl = gomock.NewController(t)
	cc.mockRepoComment = mock_comment.NewMockRepoComment(cc.ctrl)

	cc.responder = NewMockResponder(cc.mockRepoComment)

	return cc
}

func TestCreateRootComment(t *testing.T) {
	cc := prepareTestCreate(t, commentApiCreateRootCase)
	defer cc.ctrl.Finish()

	cc.mockRepoComment.
		EXPECT().
		CreateRootComment(gomock.Eq(&commentCreateCase)).
		Do(
			func(c *comment.Comment) {
				c.Id = commentApiGetRootCase.Id
				c.Path = commentApiGetRootCase.Path
			}).
		Times(1).
		Return(nil)

	cc.responder.CreateComment(cc.ctx)

	assert.Equal(t, 200, cc.ctx.Writer.Status())

	getCommentApi := api.CommentResp{}
	require.Nil(t, json.Unmarshal(cc.writer.Body.Bytes(), &getCommentApi))
	assert.Equal(t, getCommentApi, commentApiGetRootCase)
}

func TestCreateChildComment(t *testing.T) {
	cc := prepareTestCreate(t, commentApiCreateChildCase)
	defer cc.ctrl.Finish()

	cc.mockRepoComment.
		EXPECT().
		CreateChildComment(gomock.Eq(&commentCreateCase), commentApiCreateChildCase.ParentId).
		Do(
			func(c *comment.Comment, parentId int) {
				c.Id = commentApiGetChildCase.Id
				c.Path = commentApiGetChildCase.Path
			}).
		Times(1).
		Return(nil)

	cc.responder.CreateComment(cc.ctx)

	assert.Equal(t, 200, cc.ctx.Writer.Status())

	getCommentApi := api.CommentResp{}
	require.Nil(t, json.Unmarshal(cc.writer.Body.Bytes(), &getCommentApi))
	assert.Equal(t, getCommentApi, commentApiGetChildCase)
}

func TestCreateCommentFail(t *testing.T) {
	cc := prepareTestCreate(t, commentApiCreateRootCase)
	defer cc.ctrl.Finish()

	// Fail - can't create comment

	cc.mockRepoComment.
		EXPECT().
		CreateRootComment(gomock.Eq(&commentCreateCase)).
		Times(1).
		Return(errors.New(""))

	cc.responder.CreateComment(cc.ctx)

	assert.Equal(t, 500, cc.ctx.Writer.Status())

	// Fail - invalid req format

	cc.ctx.Request = httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("{")))
	cc.mockRepoComment.
		EXPECT().
		CreateRootComment(gomock.Eq(&commentCreateCase)).
		Times(0)

	cc.responder.CreateComment(cc.ctx)

	assert.Equal(t, 400, cc.ctx.Writer.Status())

	// Fail - not auth

	cc.ctx.Request = createRequest(t, commentApiCreateRootCase)
	cc.ctx.Keys = map[string]interface{}{}
	cc.mockRepoComment.
		EXPECT().
		CreateRootComment(gomock.Eq(&commentCreateCase)).
		Times(0)

	cc.responder.CreateComment(cc.ctx)

	assert.Equal(t, 401, cc.ctx.Writer.Status())
}
