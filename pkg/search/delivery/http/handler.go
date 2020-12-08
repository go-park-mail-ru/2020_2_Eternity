package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	sc "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	sc sc.SearchServiceClient
}

func NewHandler(sc sc.SearchServiceClient) *Handler {
	return &Handler{sc: sc}
}

func (h *Handler) Search(c *gin.Context) {
	sType := c.Query("type")
	cont := c.Query("content")
	if cont == "" {
		c.JSON(http.StatusBadRequest, utils.Error{Error: "no content"})
		return
	}
	last, err := strconv.Atoi(c.Query("last"))
	if err != nil {
		last = 0
	}

	switch sType {
	case "user":
		users, err := h.sc.GetUsersByName(c, &sc.UserSearch{
			Username: cont,
			Last:     int32(last),
		})

		if err != nil {
			config.Lg("deliverySearch", "users").Error(err.Error())
			c.JSON(http.StatusOK, utils.Error{Error: "Not found"})
			return
		}

		rUs := make([]domain.UserSearch, 0, len(users.Users))
		for _, u := range users.GetUsers() {
			rUs = append(rUs, domain.UserSearch{
				ID:       int(u.GetId()),
				Username: u.GetUsername(),
				Avatar:   u.GetAvatar(),
			})
		}
		c.JSON(http.StatusOK, rUs)
		return
	case "pin":
		pins, err := h.sc.GetPinsByTitle(c, &sc.PinSearch{
			Title: cont,
			Last:  int32(last),
		})
		if err != nil {
			config.Lg("deliverySearch", "Pins").Error(err.Error())
			c.JSON(http.StatusOK, utils.Error{Error: "pins not found"})
			return
		}
		rPins := make([]domain.PinResp, 0, len(pins.Pins))
		for _, p := range pins.GetPins() {
			rPins = append(rPins, domain.PinResp{
				Id:      int(p.GetId()),
				Title:   p.GetTitle(),
				Content: p.GetContent(),
				UserId:  int(p.GetUserId()),
				ImgLink: p.GetImgLink(),
			})
		}
		c.JSON(http.StatusOK, rPins)
		return
	case "board":
		boards, err := h.sc.GetBoardsByTitle(c, &sc.BoardSearch{
			Title: cont,
			Last:  int32(last),
		})
		if err != nil {
			config.Lg("deliverySearch", "Boards").Error(err.Error())
			c.JSON(http.StatusOK, utils.Error{Error: "boards not found"})
			return
		}

		rBoards := make([]domain.Board, 0, len(boards.GetBoards()))
		for _, p := range boards.GetBoards() {
			rBoards = append(rBoards, domain.Board{
				ID:       int(p.GetId()),
				Title:    p.GetTitle(),
				Content:  p.GetContent(),
				Username: p.GetUsername(),
			})
		}
		c.JSON(http.StatusOK, rBoards)
		return
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: fmt.Sprintf("Bad search type %s. Can search by user and pin", sType)})
	}
}
