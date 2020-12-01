package ws

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

type Handler struct {
	uc chat.IUsecase
	p *bluemonday.Policy
}

func NewHandler(uc chat.IUsecase, p *bluemonday.Policy) *Handler {
	return &Handler{
		uc: uc,
		p: p,
	}
}

func (h *Handler) CreateMessage(c *ws.Context) {
	userId := c.Req.UserId
	req := domainChat.CreateMessageReq{}
	if err := json.Unmarshal(c.Req.Data, &req); err != nil {
		c.AbortWithStatus(domainChat.CreateMessageRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "CreateMessage").Error("Unmarshal: ", err.Error())
		return
	}

	resp, err := h.uc.CreateMessage(&req, userId)
	if err != nil {
		c.AbortWithStatus(domainChat.CreateMessageRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "CreateMessage").Error("Usecase: ", err.Error())
		return
	}

	c.AddResponse(resp, domainChat.CreateMessageRespType, userId, http.StatusOK)
}


func (h *Handler) DeleteMessage(c *ws.Context) {
	userId := c.Req.UserId
	req := domainChat.DeleteMessageReq{}
	if err := json.Unmarshal(c.Req.Data, &req); err != nil {
		c.AbortWithStatus(domainChat.DeleteMessageRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetLastNMessages").Error("Unmarshal: ", err.Error())
		return
	}

	err := h.uc.DeleteMessage(req.MsgId)
	if err != nil {
		c.AbortWithStatus(domainChat.DeleteMessageRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetLastNMessages").Error("Usecase: ", err.Error())
		return
	}

	c.AbortWithStatus(domainChat.DeleteMessageRespType, userId, http.StatusOK)
}

func (h *Handler) GetLastNMessages(c *ws.Context) {
	userId := c.Req.UserId
	req := domainChat.GetLastNMessagesReq{}
	if err := json.Unmarshal(c.Req.Data, &req); err != nil {
		c.AbortWithStatus(domainChat.GetLastNMessagesRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetLastNMessages").Error("Unmarshal: ", err.Error())
		return
	}

	resp, err := h.uc.GetLastNMessages(&req)
	if err != nil {
		c.AbortWithStatus(domainChat.GetLastNMessagesRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetLastNMessages").Error("Usecase: ", err.Error())
		return
	}

	c.AddResponse(resp, domainChat.GetLastNMessagesRespType, userId, http.StatusOK)
}

func (h *Handler) GetNMessagesBefore(c *ws.Context) {
	userId := c.Req.UserId
	req := domainChat.GetNMessagesBeforeReq{}
	if err := json.Unmarshal(c.Req.Data, &req); err != nil {
		c.AbortWithStatus(domainChat.GetNMessagesBeforeRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetNMessagesBefore").Error("Unmarshal: ", err.Error())
		return
	}

	resp, err := h.uc.GetNMessagesBefore(&req)
	if err != nil {
		c.AbortWithStatus(domainChat.GetNMessagesBeforeRespType, userId, http.StatusBadRequest)
		config.Lg("chat_ws", "GetNMessagesBefore").Error("Usecase: ", err.Error())
		return
	}

	c.AddResponse(resp, domainChat.GetNMessagesBeforeRespType, userId, http.StatusOK)
}