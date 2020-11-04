package user

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type IUsecase interface {
	CreateUser(user *api.SignUp) (*domain.User, error)

	GetUser(id int) (*domain.User, error)
	GetUserByName(username string) (*domain.User, error)
	GetUserByNameWithFollowers(username string) (*domain.User, error)

	UpdateUser(id int, profile *api.UpdateUser) (*domain.User, error)
	UpdatePassword(id int, psswd string) error

	UpdateAvatar(id int, avatar string) error
	GetAvatar(id int) (error, string)

	Follow(following int, id int) error
	UnFollow(unfollowing int, id int) error

	GetFollowers(username string) ([]domain.User, error)
	GetFollowing(username string) ([]domain.User, error)
}
