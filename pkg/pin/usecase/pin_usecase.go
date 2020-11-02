package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"mime/multipart"
	"net/url"
	"path/filepath"
)

type Usecase struct {
	repository  pin.IRepository
	fileStorage pin.IStorage
}

func NewUsecase(r pin.IRepository, s pin.IStorage) *Usecase {
	return &Usecase{
		repository:  r,
		fileStorage: s,
	}
}

func getUrlImg(imgName string) string {
	imgUrl := url.URL{
		Scheme: config.Conf.Web.Server.Protocol,
		Host:   config.Conf.Web.Server.Host,
		Path:   filepath.Join(config.Conf.Web.Static.UrlImg, imgName),
	}

	return imgUrl.String()
}

func (u *Usecase) CreatePin(pin *domain.PinReq, file *multipart.FileHeader, userId int) (domain.PinResp, error) {
	fileName, err := utils.RandomUuid()
	if err != nil {
		config.Lg("pin_usecase", "CreatePin").Error("RandomUuid: ", err.Error())
		return domain.PinResp{}, err
	}

	if err := u.fileStorage.SaveUploadedFile(file, fileName); err != nil {
		config.Lg("pin_usecase", "CreatePin").Error("SaveUploadedFile: ", err.Error())
		return domain.PinResp{}, err
	}

	modelPin := domain.Pin{
		Title:       pin.Title,
		Content:     pin.Content,
		PictureName: fileName,
		UserId:      userId,
	}

	if err := u.repository.StorePin(&modelPin); err != nil {
		config.Lg("pin_usecase", "CreatePin").Error("StorePin: ", err.Error())
		return domain.PinResp{}, err
	}

	config.Lg("pin_usecase", "CreatePin").
		Infof(
			"Created pin {%v %v %v %v %v}",
			modelPin.Id, modelPin.Title, modelPin.Content, modelPin.PictureName, modelPin.UserId)

	return domain.PinResp{
		Id:      modelPin.Id,
		Title:   modelPin.Title,
		Content: modelPin.Content,
		ImgLink: getUrlImg(modelPin.PictureName),
		UserId:  modelPin.UserId,
	}, nil
}

func (u *Usecase) GetPinList(userId int) ([]domain.PinResp, error) {
	pins, err := u.repository.GetPinList(userId)
	if err != nil {
		config.Lg("pin_usecase", "GetPinList").Error("Repo: ", err.Error())
		return nil, err
	}

	pinsResp := []domain.PinResp{}
	for _, p := range pins {
		pinsResp = append(pinsResp, domain.PinResp{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
			ImgLink: getUrlImg(p.PictureName),
			UserId:  p.UserId,
		})
	}

	return pinsResp, nil
}
