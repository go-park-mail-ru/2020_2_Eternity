package comment

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IUsecase interface {
	CreateComment(commentReq *domain.CommentCreateReq, userId int) (domain.CommentResp, error)
	GetPinComments(pinId int) ([]domain.CommentResp, error)
	GetCommentById(commentId int) (domain.CommentResp, error)
}
