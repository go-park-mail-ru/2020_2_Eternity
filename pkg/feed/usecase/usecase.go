package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/feed"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
)

type Usecase struct {
	r feed.IRepository
}

func NewUseCase(r feed.IRepository) *Usecase {
	return &Usecase{
		r: r,
	}
}

func (uc *Usecase) GetFeed(userId int) ([]domain.PinResp, error) {
	pins, err := uc.r.GetFeed(userId)
	if err != nil {
		config.Lg("feed_usecase", "GetFeed").Error("Repo: ", err.Error())
		return nil, err
	}

	var pinsResp []domain.PinResp
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
