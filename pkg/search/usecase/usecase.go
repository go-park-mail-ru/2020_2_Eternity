package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/search"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
)

type Usecase struct {
	r search.IRepository
}

func NewUsecase(r search.IRepository) *Usecase {
	return &Usecase{r: r}
}

func (uc *Usecase) GetUsersByName(username string, last int) ([]domain.UserSearch, error) {
	return uc.r.GetUsersByName(username, last)
}

func (uc *Usecase) GetPinsByTitle(title string, last int) ([]domain.PinResp, error) {
	pins, err := uc.r.GetPinsByTitle(title, last)
	if err != nil {
		config.Lg("search_usecase", "GetPinsByTitle").Error("Repo: ", err.Error())
		return nil, err
	}

	pinsResp := make([]domain.PinResp, 0, len(pins))
	for _, p := range pins {
		pinsResp = append(pinsResp, domain.PinResp{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
			ImgLink: utils.GetUrlImg(p.PictureName),
			UserId:  p.UserId,
		})
	}

	return pinsResp, nil
}

func (uc *Usecase) GetBoardsByTitle(title string, last int) ([]domain.Board, error) {
	return uc.r.GetBoardsByTitle(title, last)
}
