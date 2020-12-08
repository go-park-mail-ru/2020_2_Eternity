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

	pUsers := make([]*sc.User, 0, len(users))
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

	pPins := make([]*sc.Pin, 0, len(pins))
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

func (s *SearchHandler) GetBoardsByTitle(ctx context.Context, boardSearch *sc.BoardSearch) (*sc.Boards, error) {
	boards, err := s.uc.GetBoardsByTitle(boardSearch.Title, int(boardSearch.Last))

	if err != nil {
		config.Lg("grpc", "GetPins").Error(err.Error())
		return &sc.Boards{}, err
	}

	pBoards := make([]*sc.Board, 0, len(boards))
	for _, b := range boards {
		pBoards = append(pBoards, &sc.Board{
			Id:       int64(b.ID),
			Title:    b.Title,
			Content:  b.Content,
			Username: b.Username,
		})
	}

	return &sc.Boards{
		Boards: pBoards,
	}, nil

}
