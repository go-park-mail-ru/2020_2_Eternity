package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	sc "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search"
)

type SearchHandler struct {
	uc search.IUsecase
}

func NewHandler(uc search.IUsecase) *SearchHandler {
	return &SearchHandler{
		uc: uc,
	}
}

func (s *SearchHandler) GetUsersByName(ctx context.Context, userSearch *sc.UserSearch) (*sc.Users, error) {
	users, err := s.uc.GetUsersByName(userSearch.Username, int(userSearch.Last))
	if err != nil {
		config.Lg("grpc", "GetUsers").Error(err.Error())
		return &sc.Users{}, err
	}

	var pUsers []*sc.User
	for _, u := range users {
		us := &sc.User{
			Username: u.Username,
			Avatar:   u.Avatar,
			Id:       int64(u.ID),
		}
		pUsers = append(pUsers, us)
	}

	return &sc.Users{
		Users: pUsers,
	}, nil
}

func (s *SearchHandler) GetPinsByTitle(ctx context.Context, pinSearch *sc.PinSearch) (*sc.Pins, error) {
	pins, err := s.uc.GetPinsByTitle(pinSearch.Title, int(pinSearch.Last))

	if err != nil {
		config.Lg("grpc", "GetPins").Error(err.Error())
		return &sc.Pins{}, err
	}

	var pPins []*sc.Pin
	for _, p := range pins {
		pPins = append(pPins, &sc.Pin{
			Id:      int64(p.Id),
			Title:   p.Title,
			Content: p.Content,
			UserId:  int32(p.UserId),
			ImgLink: p.ImgLink,
		})
	}

	return &sc.Pins{
		Pins: pPins,
	}, nil

}
