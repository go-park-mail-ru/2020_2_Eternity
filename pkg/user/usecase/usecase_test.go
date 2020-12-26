package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_user "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

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

func TestUserUsecase_GetUser(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	u := &domain.User{}
	userMockRepository.EXPECT().GetUser(1).Return(u, nil)
	if _, err := usecase.GetUser(1); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_GetUserByName(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "21savage"
	u := &domain.User{
		Username: username,
	}
	userMockRepository.EXPECT().GetUserByName(username).Return(u, nil)
	if _, err := usecase.GetUserByName(username); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_GetUserByNameWithFollowers(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "21savage"
	u := &domain.User{
		Username:  username,
		Following: 0,
		Followers: 1,
	}
	userMockRepository.EXPECT().GetUserByNameWithFollowers(username).Return(u, nil)
	if _, err := usecase.GetUserByNameWithFollowers(username); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	testUser := api.UpdateUser{
		Username: "22savage",
		Email:    "e2@mail.ru",
	}
	u := &domain.User{}
	userMockRepository.EXPECT().UpdateUser(1, &testUser).Return(u, nil)
	if _, err := usecase.UpdateUser(1, &testUser); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_UpdatePassword(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	password := "12345678"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), config.Conf.Token.Value)
	if err != nil {
		log.Fatal(err)
	}

	userMockRepository.EXPECT().UpdatePassword(1, string(hash)).Return(nil)
	if err := usecase.UpdatePassword(1, string(hash)); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_UpdateAvatar(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	filename, err := utils.RandomUuid()

	if err != nil {
		log.Fatal(err)
	}

	userMockRepository.EXPECT().UpdateAvatar(1, filename).Return(nil)
	if err := usecase.UpdateAvatar(1, filename); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_GetAvatar(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	filename, err := utils.RandomUuid()

	if err != nil {
		log.Fatal(err)
	}
	userMockRepository.EXPECT().GetAvatar(1).Return(nil, filename)
	if err, _ := usecase.GetAvatar(1); err != nil {
		assert.Equal(t, err, nil)
	}
}

//func TestUserUsecase_Follow(t *testing.T) {
//	t.Helper()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userMockRepository := mock_user.NewMockIRepository(ctrl)
//	usecase := NewUsecase(userMockRepository)
//
//	userMockRepository.EXPECT().Follow(2, 1).Return(nil)
//	if err := usecase.Follow(2, 1); err != nil {
//		assert.Equal(t, err, nil)
//	}
//}

func TestUserUsecase_UnFollow(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	userMockRepository.EXPECT().UnFollow(2, 1).Return(nil)
	if err := usecase.UnFollow(2, 1); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_GetFollowers(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "future"
	var followers []domain.User
	userMockRepository.EXPECT().GetFollowers(username).Return(followers, nil)
	if _, err := usecase.GetFollowers(username); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_GetFollowing(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "future"
	var following []domain.User
	userMockRepository.EXPECT().GetFollowing(username).Return(following, nil)
	if _, err := usecase.GetFollowing(username); err != nil {
		assert.Equal(t, err, nil)
	}
}

func TestUserUsecase_IsFollowing(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockRepository := mock_user.NewMockIRepository(ctrl)
	usecase := NewUsecase(userMockRepository)

	username := "future1"

	userMockRepository.EXPECT().IsFollowing(1, username).Return(nil)
	if err := usecase.IsFollowing(1, username); err != nil {
		assert.Equal(t, err, nil)
	}
}
