package comment

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IRepository interface {
	StoreChildComment(c *domain.Comment, parentId int) error
	StoreRootComment(c *domain.Comment) error
	GetComment(commentId int) (domain.Comment, error)
	GetPinComments(pinId int) ([]domain.Comment, error)
}
