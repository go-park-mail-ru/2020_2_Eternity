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
	req := domainChat.CreateMessageReq{}
	if err := json.Unmarshal(c.Req.Data, &req); err != nil {
		c.Status = http.StatusBadRequest
		config.Lg("chat_ws", "CreateMessage").Error("Unmarshal: ", err.Error())
		return
	}



}


