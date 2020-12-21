package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat"
	domainChat "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

const (
	ChatIdParam = "ch_id"
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



func (h *Handler) CreateChat(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("chat_http", "CreateChat").Error("Can't get claims")
		return
	}

	req := domainChat.ChatCreateReq{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "CreateChat").Error("BindJSON ", err.Error())
		return
	}

	req.CollocutorName = h.p.Sanitize(req.CollocutorName)

	resp, err := h.uc.CreateChat(&req, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "CreateChat").Error("uc.CreateChat ", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
	//c.Set(domain.NotificationKey, domain.NoteChat{
	//	Id: resp.Id,
	//	CreatorId: userId,
	//	CreationTime: resp.CreationTime,
	//	LastMsgContent: resp.LastMsgContent,
	//	LastMsgUsername: resp.LastMsgUsername,
	//	LastMsgTime: resp.LastMsgTime,
	//	CollocutorName: resp.CollocutorName,
	//	CollocutorAvatarLink: resp.CollocutorAvatarLink,
	//	NewMessages: resp.NewMessages,
	//})
}

func (h *Handler) GetChatById(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("chat_http", "GetChatById").Error("Can't get claims")
		return
	}

	chatId, err := utils.GetIntParam(c, ChatIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "GetChatById").Error("Can't get param")
		return
	}

	resp, err := h.uc.GetChatById(chatId, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "GetChatById").Error("Usecase ", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserChats(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("chat_http", "GetUserChats").Error("Can't get claims")
		return
	}

	resp, err := h.uc.GetUserChats(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "GetUserChats").Error("Usecase ", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) MarkAllMessagesRead(c *gin.Context) {
	userId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("chat_http", "MarkAllMessagesRead").Error("Can't get claims")
		return
	}

	req := domainChat.MarkMessagesReadReq{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "MarkAllMessagesRead").Error("BindJSON ", err.Error())
		return
	}

	err := h.uc.MarkAllMessagesRead(req.ChatId, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "MarkAllMessagesRead").Error("Usecase ", err.Error())
		return
	}

	c.Status(http.StatusOK)
}




func ServeWs(s ws.IServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, ok := jwthelper.GetClaims(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			config.Lg("pin_http", "ServeWs").Error("Can't get claims")
			return
		}

		//userId := 2 // Note: for tests


		if err := s.RegisterClient(c.Writer, c.Request, userId); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			config.Lg("chat_http", "ServeWs").Error(err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}
