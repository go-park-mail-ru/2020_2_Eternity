package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
)

const (
	PinIdParam     = "pin_id"
	CommentIdParam = "comment_id"
)

/*
	Get   pin/comments/:comm_id	 - get concrete comment
	Get   pin/:pin_id/comments   - get all comments for pin
	Post  pin/:pin_id/comments   - create new comment for pin
*/

type ResponderComment struct {
	RepoComment RepoComment
}

func NewResponder() *ResponderComment {
	return &ResponderComment{
		RepoComment: &RepoCommentInstance{},
	}
}

func (rc *ResponderComment) CreateComment(c *gin.Context) {
	claimsId, ok := user.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("comment", "CreateComment").Error("Can't get claims")
		return
	}

	createCommentApi := api.CreateComment{}
	if err := c.BindJSON(&createCommentApi); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "CreateComment").Error("BindJSON ", err.Error())
		return
	}

	comment := Comment{
		Content: createCommentApi.Content,
		PinId:   createCommentApi.PinId,
		UserId:  claimsId,
	}

	var err error
	if createCommentApi.IsRoot {
		err = rc.RepoComment.CreateRootComment(&comment)
	} else {
		err = rc.RepoComment.CreateChildComment(&comment, createCommentApi.ParentId)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		config.Lg("comment", "CreateComment").Error(err.Error())
		return
	}

	getCommentApi := api.GetComment{
		Id:      comment.Id,
		Path:    comment.Path,
		Content: comment.Content,
		PinId:   comment.PinId,
		UserId:  comment.UserId,
	}

	c.JSON(http.StatusOK, getCommentApi)
}

func (rc *ResponderComment) GetComments(c *gin.Context) {
	pinId, err := utils.GetIntParam(c, PinIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetComments").Error(err.Error())
		return
	}

	comments, err := rc.RepoComment.GetAllComments(pinId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetComments").Error(err.Error())
		return
	}

	commentsApi := []api.GetComment{}
	for _, comment := range comments {
		commentsApi = append(commentsApi, api.GetComment{
			Id:      comment.Id,
			Path:    comment.Path,
			Content: comment.Content,
			PinId:   comment.PinId,
			UserId:  comment.UserId,
		})
	}

	c.JSON(http.StatusOK, commentsApi)
}

func (rc *ResponderComment) GetCommentById(c *gin.Context) {
	commentId, err := utils.GetIntParam(c, CommentIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetCommentById").Error(err.Error())
		return
	}

	comment, err := rc.RepoComment.GetComment(commentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetCommentById").Error(err.Error())
		return
	}

	commentApi := api.GetComment{
		Id:      comment.Id,
		Path:    comment.Path,
		Content: comment.Content,
		PinId:   comment.PinId,
		UserId:  comment.UserId,
	}

	c.JSON(http.StatusOK, commentApi)
}
