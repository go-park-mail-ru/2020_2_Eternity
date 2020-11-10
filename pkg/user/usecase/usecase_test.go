package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_user "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	testUser := api.SignUp{
		Username: "21savage",
		Password: "12345678",
		Email:    "e@mail.ru",
	}
	u := &domain.User{}
	userMockRepository.EXPECT().CreateUser(&testUser).Return(u, nil)
	if _, err := usecase.CreateUser(&testUser); err != nil {
		assert.Equal(t, err, nil)
	}
}
