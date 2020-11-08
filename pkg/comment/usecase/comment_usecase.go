package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type Usecase struct {
	repository  comment.IRepository
}

func NewUsecase(r comment.IRepository) *Usecase {
	return &Usecase{
		repository:  r,
	}
}


func (uc *Usecase) CreateComment(commentReq *domain.CommentCreateReq, userId int) (domain.CommentResp, error) {
	modelComment := domain.Comment{
		Content: commentReq.Content,
		PinId: commentReq.PinId,
		UserId: userId,
	}

	var err error
	if commentReq.IsRoot {
		err = uc.repository.StoreRootComment(&modelComment)
	} else {
		err = uc.repository.StoreChildComment(&modelComment, commentReq.ParentId)
	}

	if err != nil {
		config.Lg("comment_usecase", "CreateComment").Error(err.Error())
		return domain.CommentResp{}, err
	}

	return domain.CommentResp{
		Id: modelComment.Id,
		Path : modelComment.Path,
		Content: modelComment.Content,
		PinId: modelComment.PinId,
		UserId: modelComment.UserId,
	}, nil
}


func (uc *Usecase) GetPinComments(pinId int) ([]domain.CommentResp, error) {
	modelComments, err := uc.repository.GetPinComments(pinId)
	if err != nil {
		config.Lg("comment_usecase", "GetPinComments").Error(err.Error())
		return nil, err
	}

	commentsResp := []domain.CommentResp{}
	for _, mComment := range modelComments {
		commentsResp = append(commentsResp, domain.CommentResp{
			Id: mComment.Id,
			Path : mComment.Path,
			Content: mComment.Content,
			PinId: mComment.PinId,
			UserId: mComment.UserId,
		})
	}

	return commentsResp, nil
}


func (uc *Usecase) GetCommentById(commentId int) (domain.CommentResp, error) {
	modelComment, err := uc.repository.GetComment(commentId)
	if err != nil {
		config.Lg("comment_usecase", "GetCommentById").Error(err.Error())
		return domain.CommentResp{}, err
	}

	return domain.CommentResp{
		Id: modelComment.Id,
		Path : modelComment.Path,
		Content: modelComment.Content,
		PinId: modelComment.PinId,
		UserId: modelComment.UserId,
	}, nil
}