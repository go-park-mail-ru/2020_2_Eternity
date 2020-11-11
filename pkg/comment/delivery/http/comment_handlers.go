package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

const (
	PinIdParam     = "pin_id"
	CommentIdParam = "comment_id"
)

type Handler struct {
	uc comment.IUsecase
	p  *bluemonday.Policy
}

func NewHandler(uc comment.IUsecase, p *bluemonday.Policy) *Handler {
	return &Handler{
		uc: uc,
		p:  p,
	}
}

func (h *Handler) CreateComment(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("comment_http", "CreateComment").Error("Can't get claims")
		return
	}

	commentReq := domain.CommentCreateReq{}
	if err := c.BindJSON(&commentReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment_http", "CreateComment").Error("BindJSON ", err.Error())
		return
	}

	if err := commentReq.Validate(); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment_http", "CreateComment").Error("Validate: ", err.Error())
		return
	}

	commentReq.Content = h.p.Sanitize(commentReq.Content)

	commentResp, err := h.uc.CreateComment(&commentReq, userId)
	if err != nil {
		// TODO (Paul S) Error switch
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment_http", "CreateComment").Error("uc.CreateComment ", err.Error())
		return
	}

	c.JSON(http.StatusOK, commentResp)
	c.Set(domain.NotificationKey, domain.NoteComment{
		Id:      commentResp.Id,
		Path:    commentResp.Path,
		Content: commentResp.Content,
		PinId:   commentResp.PinId,
		UserId:  commentResp.UserId,
	})
}

func (h *Handler) GetPinComments(c *gin.Context) {
	pinId, err := utils.GetIntParam(c, PinIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment_http", "GetPinComments").Error(err.Error())
		return
	}

	commentsResp, err := h.uc.GetPinComments(pinId)
	if err != nil {
		// TODO (Paul S) Error switch
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment_http", "GetPinComments").Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, commentsResp)
}

func (h *Handler) GetCommentById(c *gin.Context) {
	commentId, err := utils.GetIntParam(c, CommentIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetCommentById").Error(err.Error())
		return
	}

	commentResp, err := h.uc.GetCommentById(commentId)
	if err != nil {
		// TODO (Paul S) Error switch
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("comment", "GetCommentById").Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, commentResp)
}
