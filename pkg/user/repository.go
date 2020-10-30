package user

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/models"
)

type IRepository interface {
	CreateUser(user *api.SignUp) (*models.User, error)

	GetUser(id int) (*models.User, error)
	GetUserByName(username string) (*models.User, error)
	GetUserByNameWithFollowers(username string) (*models.User, error)

	UpdateUser(id int, profile *api.UpdateUser) (*models.User, error)
	UpdatePassword(id int, psswd string) error

	UpdateAvatar(id int, avatar string) error
	GetAvatar(id int) (error, string)

	Follow(following int, id int) error
	UnFollow(unfollowing int, id int) error
}
