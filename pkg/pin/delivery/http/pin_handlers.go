package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Handler struct {
	uc pin.IUsecase
	p  *bluemonday.Policy
}

func NewHandler(uc pin.IUsecase, p *bluemonday.Policy) *Handler {
	return &Handler{
		uc: uc,
		p:  p,
	}
}

type FormCreatePin struct {
	CreatePinReq *domain.PinReq        `form:"data" binding:"required"`
	Avatar       *multipart.FileHeader `form:"img" binding:"required"`
}

func (h *Handler) CreatePin(c *gin.Context) {
	userId, ok := auth.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("pin_http", "CreatePin").Error("Can't get claims")
		return
	}

	formPin := FormCreatePin{}
	if err := c.Bind(&formPin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin_http", "CreatePin").Error("Bind: ", err.Error())
		return
	}

	if err := formPin.CreatePinReq.Validate(); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin_http", "CreatePin").Error("Validate: ", err.Error())
		return
	}

	h.sanitize(formPin.CreatePinReq)

	pinResp, err := h.uc.CreatePin(formPin.CreatePinReq, formPin.Avatar, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		config.Lg("pin_http", "CreatePin").Error("uc.CreatePin: ", err.Error())
		return
	}

	c.JSON(http.StatusOK, pinResp)
	c.Set(domain.NotificationKey, domain.NotePin{
		Id: pinResp.Id,
		Title: pinResp.Title,
		ImgLink: pinResp.ImgLink,
		UserId: pinResp.UserId,
	})
}

func (h *Handler) GetPin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "not integer id"})
		return
	}
	p, err := h.uc.GetPin(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.Error{Error: "not found pin or fake id"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) GetAllPins(c *gin.Context) {
	username := h.p.Sanitize(c.Param("username"))

	pins, err := h.uc.GetPinList(username)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin_http", "GetAllPins").Error("uc.GetPinList: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, pins)
}

func (h *Handler) GetPinsFromBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin_http", "GetPinsFromBoard").Error("Bind: ", "No param")
		return
	}

	pins, err := h.uc.GetPinBoardList(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "Cant show pins from board"})
		config.Lg("pin_http", "GetPinsFromBoard").Error("Get pins", err.Error())
		return
	}
	c.JSON(http.StatusOK, pins)
}

func (h *Handler) sanitize(f *domain.PinReq) {
	f.Title = h.p.Sanitize(f.Title)
	f.Content = h.p.Sanitize(f.Content)
}
