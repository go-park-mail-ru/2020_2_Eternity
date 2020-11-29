package ws

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
	"github.com/microcosm-cc/bluemonday"
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

func (h *Handler) CreateMessage (c *ws.Context) {
	return
}


