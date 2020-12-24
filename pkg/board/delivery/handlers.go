package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/board"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strconv"
)

type Handler struct {
	uc board.IUsecase
	p  *bluemonday.Policy
}

func NewHandler(uc board.IUsecase, p *bluemonday.Policy) *Handler {
	return &Handler{
		uc: uc,
		p:  p,
	}
}

func (h *Handler) CreateBoard(c *gin.Context) {
	claimsId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	b := api.CreateBoard{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	if err := h.sanitize(&b); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: err.Error()})
		return
	}

	rb, err := h.uc.CreateBoard(claimsId, &b)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, *rb)
}

func (h *Handler) GetBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "not integer id"})
		return
	}
	b, err := h.uc.GetBoard(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.Error{Error: "not found board or fake id"})
		return
	}
	c.JSON(http.StatusOK, *b)
}

func (h *Handler) GetAllBoardsbyUser(c *gin.Context) {
	username := h.p.Sanitize(c.Param("username"))
	b, err := h.uc.GetAllBoardsByUser(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, b)
}

func (h *Handler) AttachPinToBoard(c *gin.Context) {
	bp, status, err := h.prepAtDet(c)
	if err != nil {
		c.AbortWithStatusJSON(status, *err)
		return
	}
	if err := h.uc.AttachPin(bp.BoardID, bp.PinID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "Cannot attach"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DetachPinFromBoard(c *gin.Context) {
	bp, status, err := h.prepAtDet(c)
	if err != nil {
		c.AbortWithStatusJSON(status, *err)
		return
	}
	if err := h.uc.DetachPin(bp.BoardID, bp.PinID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "Cannot detach"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) prepAtDet(c *gin.Context) (*api.AttachDetachPin, int, *utils.Error) {
	claimsId, ok := jwthelper.GetClaims(c)
	if !ok {
		return nil, http.StatusUnauthorized, &utils.Error{Error: "invalid token"}
	}

	bp := api.AttachDetachPin{}
	if err := c.BindJSON(&bp); err != nil {
		return nil, http.StatusBadRequest, &utils.Error{Error: "[BindJSON]: " + err.Error()}
	}
	if err := h.uc.CheckOwner(claimsId, bp.BoardID); err != nil {
		return &bp, http.StatusBadRequest, &utils.Error{Error: err.Error()}
	}
	return &bp, http.StatusOK, nil
}

func (h *Handler) GetBoardsPinNotAttach(c *gin.Context) {
	claimsId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	p := c.Param("pin_id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "not id"})
		return
	}

	boards, err := h.uc.GetBoardsPinNotAttach(claimsId, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "id not found"})
		return
	}
	c.JSON(http.StatusOK, boards)
}

func (h *Handler) sanitize(b *api.CreateBoard) error {
	b.Content = h.p.Sanitize(b.Content)
	b.Title = h.p.Sanitize(b.Title)
	if b.Content == "" && b.Title == "" {
		return errors.New("empty board")
	}
	return nil
}
