package repository

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/models"
)

type Repository struct {
	db database.DBInterface
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
	}
}

func (r *Repository) CreateUser(user *api.SignUp) (*models.User, error) {
	return r.db.CreateUser(user)
}

func (r *Repository) GetUser(id int) (*models.User, error) {
	return r.db.GetUser(id)
}

func (r *Repository) GetUserByName(username string) (*models.User, error) {
	return r.db.GetUserByName(username)
}

func (r *Repository) GetUserByNameWithFollowers(username string) (*models.User, error) {
	return r.db.GetUserByName(username)
}

func (r *Repository) UpdateUser(id int, profile *api.UpdateUser) (*models.User, error) {
	return r.db.UpdateUser(id, profile)
}
func (r *Repository) UpdatePassword(id int, psswd string) error {
	return r.db.UpdatePassword(id, psswd)
}

func (r *Repository) UpdateAvatar(id int, avatar string) error {
	return r.db.UpdateAvatar(id, avatar)
}
func (r *Repository) GetAvatar(id int) (error, string) {
	return r.db.GetAvatar(id)
}

func (r *Repository) Follow(following int, id int) error {
	return r.db.Follow(following, id)
}
func (r *Repository) UnFollow(unfollowing int, id int) error {
	return r.db.UnFollow(unfollowing, id)
}
