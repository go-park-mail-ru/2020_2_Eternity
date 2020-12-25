package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
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

func (uc *UserUsecase) CreateUser(user *api.SignUp) (*domain.User, error) {
	return uc.repo.CreateUser(user)
}

func (uc *UserUsecase) GetUser(id int) (*domain.User, error) {
	return uc.repo.GetUser(id)
}

func (uc *UserUsecase) GetUserByName(username string) (*domain.User, error) {
	return uc.repo.GetUserByName(username)
}

func (uc *UserUsecase) GetUserByNameWithFollowers(username string) (*domain.User, error) {
	return uc.repo.GetUserByNameWithFollowers(username)
}

func (uc *UserUsecase) UpdateUser(id int, profile *api.UpdateUser) (*domain.User, error) {
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

func (uc *UserUsecase) Follow(following int, id int) (string, error) {
	return uc.repo.Follow(following, id)
}
func (uc *UserUsecase) UnFollow(unfollowing int, id int) error {
	return uc.repo.UnFollow(unfollowing, id)
}

func (uc *UserUsecase) GetFollowers(username string) ([]domain.User, error) {
	return uc.repo.GetFollowers(username)
}
func (uc *UserUsecase) GetFollowing(username string) ([]domain.User, error) {
	return uc.repo.GetFollowing(username)
}

func (uc *UserUsecase) IsFollowing(id int, username string) error {
	return uc.repo.IsFollowing(id, username)
}

func (uc *UserUsecase) GetPopularUsers(limit int) ([]domain.UserSearch, error) {
	return uc.repo.GetPopularUsers(limit)
}
