package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
)

type UserUsecase struct {
	repo user.IRepository
}

func NewUsecase(repo user.IRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) CreateUser(user *api.SignUp) (*models.User, error) {
	return uc.repo.CreateUser(user)
}

func (uc *UserUsecase) GetUser(id int) (*models.User, error) {
	return uc.repo.GetUser(id)
}

func (uc *UserUsecase) GetUserByName(username string) (*models.User, error) {
	return uc.repo.GetUserByName(username)
}

func (uc *UserUsecase) GetUserByNameWithFollowers(username string) (*models.User, error) {
	return uc.repo.GetUserByName(username)
}

func (uc *UserUsecase) UpdateUser(id int, profile *api.UpdateUser) (*models.User, error) {
	return uc.repo.UpdateUser(id, profile)
}
func (uc *UserUsecase) UpdatePassword(id int, psswd string) error {
	return uc.repo.UpdatePassword(id, psswd)
}

func (uc *UserUsecase) UpdateAvatar(id int, avatar string) error {
	return uc.repo.UpdateAvatar(id, avatar)
}
func (uc *UserUsecase) GetAvatar(id int) (error, string) {
	return uc.repo.GetAvatar(id)
}

func (uc *UserUsecase) Follow(following int, id int) error {
	return uc.repo.Follow(following, id)
}
func (uc *UserUsecase) UnFollow(unfollowing int, id int) error {
	return uc.repo.UnFollow(unfollowing, id)
}
