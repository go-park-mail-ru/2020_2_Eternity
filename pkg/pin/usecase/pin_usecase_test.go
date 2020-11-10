package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	pinUsecase "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/usecase"
	mock_pin "github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/mock"
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

	pinReq = domain.PinReq{
		Title:       "title",
		Content:     "content",
	}

	pinResp = domain.PinResp{
		Id: 		 pinId,
		Title:       pinReq.Title,
		Content:     pinReq.Content,
		ImgLink: 	 "link",
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

	code := m.Run()
	os.Exit(code)
}

func TestCreatePin(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockRepo := mock_pin.NewMockIRepository(mockCtr)
	mockStorage := mock_pin.NewMockIStorage(mockCtr)

	uc := pinUsecase.NewUsecase(mockRepo, mockStorage)

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
