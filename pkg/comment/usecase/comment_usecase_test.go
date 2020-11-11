package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_comment "github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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

	commentReqChild = domain.CommentCreateReq{
		IsRoot:   false,
		ParentId: 1,
		Content:  "content",
		PinId:    2,
	}

	commentRespChild = domain.CommentResp{
		Id:      2,
		Path:    []int32{1, 2},
		Content: commentReqChild.Content,
		PinId:   commentReqChild.PinId,
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
	config.Conf = config.NewTestConfig()
	return true
}()

func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

func TestCreateRootComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_comment.NewMockIRepository(ctrl)

	uc := NewUsecase(mockRepo)

	// Success

	mockRepo.EXPECT().
		StoreRootComment(gomock.Eq(&domain.Comment{
			Content: commentReqRoot.Content,
			PinId:   commentReqRoot.PinId,
			UserId:  userId,
		})).
		DoAndReturn(func(c *domain.Comment) error {
			c.Id = commentRespRoot.Id
			c.Path = commentRespRoot.Path
			return nil
		}).
		Times(1)

	cResp, err := uc.CreateComment(&commentReqRoot, userId)
	assert.Nil(t, err)
	assert.Equal(t, commentRespRoot, cResp)

	// Fail

	mockRepo.EXPECT().
		StoreRootComment(gomock.Eq(&domain.Comment{
			Content: commentReqRoot.Content,
			PinId:   commentReqRoot.PinId,
			UserId:  userId,
		})).
		Return(errors.New("")).
		Times(1)

	_, err = uc.CreateComment(&commentReqRoot, userId)
	assert.NotNil(t, err)
}

func TestCreateChildComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_comment.NewMockIRepository(ctrl)

	uc := NewUsecase(mockRepo)

	// Success

	mockRepo.EXPECT().
		StoreChildComment(gomock.Eq(&domain.Comment{
			Content: commentReqChild.Content,
			PinId:   commentReqChild.PinId,
			UserId:  userId,
		}), gomock.Eq(commentReqChild.ParentId)).
		DoAndReturn(func(c *domain.Comment, parentId int) error {
			c.Id = commentRespChild.Id
			c.Path = commentRespChild.Path
			return nil
		}).
		Times(1)

	cResp, err := uc.CreateComment(&commentReqChild, userId)
	assert.Nil(t, err)
	assert.Equal(t, commentRespChild, cResp)
}

func TestGetPinComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_comment.NewMockIRepository(ctrl)

	uc := NewUsecase(mockRepo)

	// Success

	mockRepo.EXPECT().
		GetPinComments(gomock.Eq(pinId)).
		DoAndReturn(func(pinId int) ([]domain.Comment, error) {
			commentsResp := []domain.Comment{}
			for _, mComment := range commentRespMany {
				commentsResp = append(commentsResp, domain.Comment{
					Id:      mComment.Id,
					Path:    mComment.Path,
					Content: mComment.Content,
					PinId:   mComment.PinId,
					UserId:  mComment.UserId,
				})
			}

			return commentsResp, nil
		}).
		Times(1)

	cResp, err := uc.GetPinComments(pinId)
	assert.Nil(t, err)
	assert.Equal(t, commentRespMany, cResp)

	// Fail

	mockRepo.EXPECT().
		GetPinComments(gomock.Eq(pinId)).
		Return(nil, errors.New("")).
		Times(1)

	_, err = uc.GetPinComments(pinId)
	assert.NotNil(t, err)
}

func TestGetCommentById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_comment.NewMockIRepository(ctrl)

	uc := NewUsecase(mockRepo)

	// Success

	mockRepo.EXPECT().
		GetComment(gomock.Eq(commentId)).
		Return(domain.Comment{
			Id:      commentRespOne.Id,
			Path:    commentRespOne.Path,
			Content: commentRespOne.Content,
			PinId:   commentRespOne.PinId,
			UserId:  commentRespOne.UserId,
		}, nil).
		Times(1)

	cResp, err := uc.GetCommentById(commentId)
	assert.Nil(t, err)
	assert.Equal(t, commentRespOne, cResp)

	// Fail

	mockRepo.EXPECT().
		GetComment(gomock.Eq(commentId)).
		Return(domain.Comment{}, errors.New("")).
		Times(1)

	_, err = uc.GetCommentById(commentId)
	assert.NotNil(t, err)
}
