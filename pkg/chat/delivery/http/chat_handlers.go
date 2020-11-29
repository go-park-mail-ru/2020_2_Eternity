package http

import (
	"github.com/gin-gonic/gin"
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



func (h *Handler) CreateChat(c *gin.Context) {
	//_, ok := jwthelper.GetClaims(c)
	//if !ok {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//	config.Lg("chat_http", "CreateChat").Error("Can't get claims")
	//	return
	//}

	req := domainChat.ChatCreateReq{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "CreateChat").Error("BindJSON ", err.Error())
		return
	}

	req.CollocutorName = h.p.Sanitize(req.CollocutorName)

	resp, err := h.uc.CreateChat(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("chat_http", "CreateChat").Error("uc.CreateChat ", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ServeWs(s ws.IServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		//userId, ok := jwthelper.GetClaims(c)
		//if !ok {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	config.Lg("pin_http", "ServeWs").Error("Can't get claims")
		//	return
		//}

		userId := 2 // Note: for tests

		if err := s.RegisterClient(c.Writer, c.Request, userId); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			config.Lg("chat_http", "ServeWs").Error(err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}
