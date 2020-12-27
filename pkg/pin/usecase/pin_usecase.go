package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"log"
	"mime/multipart"
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

func (u *Usecase) CreatePin(pin *domain.PinReq, file *multipart.FileHeader, userId int) (domain.PinResp, error) {
	fileName, err := utils.RandomUuid()
	if err != nil {
		config.Lg("pin_usecase", "CreatePin").Error("RandomUuid: ", err.Error())
		return domain.PinResp{}, err
	}

	err = u.fileStorage.SaveUploadedFile(file, &fileName)
	if err != nil {
		config.Lg("pin_usecase", "CreatePin").Error("SaveUploadedFile: ", err.Error())
		return domain.PinResp{}, err
	}

	h, w, err := u.fileStorage.GetImageSizes(fileName)
	if err != nil {
		config.Lg("pin_usecase", "CreatePin").Error(err.Error())
		return domain.PinResp{}, err
	}

	log.Println(h, w)

	modelPin := domain.Pin{
		Title:         pin.Title,
		Content:       pin.Content,
		PictureName:   fileName,
		PictureHeight: h,
		PictureWidth:  w,
		UserId:        userId,
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
		Id:            modelPin.Id,
		Title:         modelPin.Title,
		Content:       modelPin.Content,
		PictureHeight: modelPin.PictureHeight,
		PictureWidth:  modelPin.PictureWidth,
		ImgLink:       utils.GetUrlImg(modelPin.PictureName),
		UserId:        modelPin.UserId,
	}, nil
}

func (u *Usecase) GetPin(id int) (domain.PinResp, error) {
	modelPin, err := u.repository.GetPin(id)

	if err != nil {
		config.Lg("pin_usecase", "GetPin").Error("Repo: ", err.Error())
		return domain.PinResp{}, err
	}

	return domain.PinResp{
		Id:            modelPin.Id,
		Title:         modelPin.Title,
		Content:       modelPin.Content,
		PictureHeight: modelPin.PictureHeight,
		PictureWidth:  modelPin.PictureWidth,
		ImgLink:       utils.GetUrlImg(modelPin.PictureName),
		UserId:        modelPin.UserId,
		Username:      modelPin.Username,
		Avatar:        modelPin.Avatar,
	}, nil
}

func (u *Usecase) GetPinList(username string) ([]domain.PinResp, error) {
	pins, err := u.repository.GetPinList(username)
	if err != nil {
		config.Lg("pin_usecase", "GetPinList").Error("Repo: ", err.Error())
		return nil, err
	}

	pinsResp := make([]domain.PinResp, 0, len(pins))
	for _, p := range pins {
		pinsResp = append(pinsResp, domain.PinResp{
			Id:            p.Id,
			Title:         p.Title,
			Content:       p.Content,
			PictureHeight: p.PictureHeight,
			PictureWidth:  p.PictureWidth,
			ImgLink:       utils.GetUrlImg(p.PictureName),
			UserId:        p.UserId,
		})
	}

	return pinsResp, nil
}

func (u *Usecase) GetPinBoardList(boardId int, limit int) ([]domain.PinResp, error) {
	pins, err := u.repository.GetPinBoardList(boardId, limit)
	if err != nil {
		config.Lg("pin_usecase", "GetPinBoardList").Error("Repo: ", err.Error())
		return nil, err
	}

	pinsResp := make([]domain.PinResp, 0, len(pins))
	for _, p := range pins {
		pinsResp = append(pinsResp, domain.PinResp{
			Id:            p.Id,
			Title:         p.Title,
			Content:       p.Content,
			PictureHeight: p.PictureHeight,
			PictureWidth:  p.PictureWidth,
			ImgLink:       utils.GetUrlImg(p.PictureName),
			UserId:        p.UserId,
		})
	}

	return pinsResp, nil
}
