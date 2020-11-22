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
	/*claimsId, ok := jwthelper.GetClaims(c)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}*/

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

		var rUs []domain.UserSearch
		for _, u := range users.Users {
			rUs = append(rUs, domain.UserSearch{
				ID:       int(u.Id),
				Username: u.Username,
				Avatar:   u.Avatar,
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
		var rPins []domain.PinResp
		for _, p := range pins.Pins {
			rPins = append(rPins, domain.PinResp{
				Id:      int(p.Id),
				Title:   p.Title,
				Content: p.Content,
				UserId:  int(p.UserId),
				ImgLink: p.ImgLink,
			})
		}
		c.JSON(http.StatusOK, rPins)
		return
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: fmt.Sprintf("Bad search type %s. Can search by user and pin", sType)})
	}

}
