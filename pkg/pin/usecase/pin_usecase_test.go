package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_pin "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"os"
	"testing"
)



var (

	userId = 3
	pinId = 4
	username = "username123"
	boardId = 6
	pictureName = "picture"

	pinReq = domain.PinReq{
		Title:       "title",
		Content:     "content",
	}

	pinResp = domain.PinResp{
		Id: 		 pinId,
		Title:       pinReq.Title,
		Content:     pinReq.Content,
		ImgLink: 	 "utils.GetUrlImg(pictureName)",
		UserId:      userId,
	}

	pinRespMany = []domain.PinResp{
		{
			Id: 		 3,
			Title:       "tittle123",
			Content:     "content14",
			ImgLink: 	 "link4213",
			UserId:      userId,
		},
		{
			Id: 		 4,
			Title:       "tittle13",
			Content:     "content4",
			ImgLink: 	 "link13",
			UserId:      userId,
		},
	}
)

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	pinResp.ImgLink = utils.GetUrlImg(pictureName)
	pinRespMany[0].ImgLink = utils.GetUrlImg(pictureName)
	pinRespMany[1].ImgLink = utils.GetUrlImg(pictureName)
	code := m.Run()
	os.Exit(code)
}

func TestCreatePin(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockRepo := mock_pin.NewMockIRepository(mockCtr)
	mockStorage := mock_pin.NewMockIStorage(mockCtr)

	uc := NewUsecase(mockRepo, mockStorage)

	mockRepo.EXPECT().
		StorePin(gomock.Any()).
		DoAndReturn(func (p *domain.Pin) error {
			p.Id = pinId
			return nil
		}).
		Times(1)

	mockStorage.EXPECT().
		SaveUploadedFile(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	pResp, err := uc.CreatePin(&pinReq, &multipart.FileHeader{}, userId)
	assert.Nil(t, err)
	assert.Equal(t, pinResp.Id, pResp.Id)
	assert.Equal(t, pinResp.Title, pResp.Title)
	assert.Equal(t, pinResp.Content, pResp.Content)
	assert.Equal(t, pinResp.UserId, pResp.UserId)

	// fail storage

	mockStorage.EXPECT().
		SaveUploadedFile(gomock.Any(), gomock.Any()).
		Return(errors.New("")).
		Times(1)

	mockRepo.EXPECT().
		StorePin(gomock.Any()).
		Return(errors.New("")).
		Times(0)


	_, err = uc.CreatePin(&pinReq, &multipart.FileHeader{}, userId)
	assert.NotNil(t, err)

	// fail repo

	mockStorage.EXPECT().
		SaveUploadedFile(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	mockRepo.EXPECT().
		StorePin(gomock.Any()).
		Return(errors.New("")).
		Times(1)


	_, err = uc.CreatePin(&pinReq, &multipart.FileHeader{}, userId)
	assert.NotNil(t, err)
}


func TestGetPin(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockRepo := mock_pin.NewMockIRepository(mockCtr)
	mockStorage := mock_pin.NewMockIStorage(mockCtr)

	uc := NewUsecase(mockRepo, mockStorage)

	// Success

	mockRepo.EXPECT().
		GetPin(gomock.Eq(pinId)).
		Return(domain.Pin{
			Id: 		 pinResp.Id,
			Title:       pinResp.Title,
			Content:     pinResp.Content,
			PictureName: pictureName,
			UserId:      pinResp.UserId,
		}, nil).
		Times(1)

	pResp, err := uc.GetPin(pinId)

	assert.Nil(t, err)
	assert.Equal(t, pinResp.Id, pResp.Id)
	assert.Equal(t, pinResp.Title, pResp.Title)
	assert.Equal(t, pinResp.Content, pResp.Content)
	assert.Equal(t, pinResp.UserId, pResp.UserId)

	// Fail

	mockRepo.EXPECT().
		GetPin(gomock.Eq(pinId)).
		Return(domain.Pin{}, errors.New("")).
		Times(1)

	_, err = uc.GetPin(pinId)
	assert.NotNil(t, err)


}


func TestGetPinList(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockRepo := mock_pin.NewMockIRepository(mockCtr)
	mockStorage := mock_pin.NewMockIStorage(mockCtr)

	uc := NewUsecase(mockRepo, mockStorage)

	// Success

	mockRepo.EXPECT().
		GetPinList(gomock.Eq(username)).
		DoAndReturn(func (u string) ([]domain.Pin, error) {
			ps := []domain.Pin{}
			for _, p := range pinRespMany {
				ps = append(ps, domain.Pin{
					Id:          p.Id,
					Title:       p.Title,
					Content:     p.Content,
					PictureName: pictureName,
					UserId:      p.UserId,
				})
			}
			return ps, nil
		}).
		Times(1)

	psResp, err := uc.GetPinList(username)

	assert.Nil(t, err)
	assert.Equal(t, pinRespMany, psResp)


	// Fail

	mockRepo.EXPECT().
		GetPinList(gomock.Eq(username)).
		Return([]domain.Pin{}, errors.New("")).
		Times(1)

	_, err = uc.GetPinList(username)
	assert.NotNil(t, err)
}


func TestGetPinBoardList(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockRepo := mock_pin.NewMockIRepository(mockCtr)
	mockStorage := mock_pin.NewMockIStorage(mockCtr)

	uc := NewUsecase(mockRepo, mockStorage)

	// Success

	mockRepo.EXPECT().
		GetPinBoardList(gomock.Eq(boardId)).
		DoAndReturn(func (b int) ([]domain.Pin, error) {
			ps := []domain.Pin{}
			for _, p := range pinRespMany {
				ps = append(ps, domain.Pin{
					Id:          p.Id,
					Title:       p.Title,
					Content:     p.Content,
					PictureName: pictureName,
					UserId:      p.UserId,
				})
			}
			return ps, nil
		}).
		Times(1)

	psResp, err := uc.GetPinBoardList(boardId)

	assert.Nil(t, err)
	assert.Equal(t, pinRespMany, psResp)


	// Fail

	mockRepo.EXPECT().
		GetPinBoardList(gomock.Eq(boardId)).
		Return([]domain.Pin{}, errors.New("")).
		Times(1)

	_, err = uc.GetPinBoardList(boardId)
	assert.NotNil(t, err)
}