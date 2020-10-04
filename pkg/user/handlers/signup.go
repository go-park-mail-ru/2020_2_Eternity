package handlers

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Handler struct {
	users model.IUsers
}

func NewHandler(users model.IUsers) *Handler {
	return &Handler{users: users}
}

func ValidProfile(profile api.SignUp) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Email, validation.Required, is.EmailFormat),
		validation.Field(&profile.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric),
		validation.Field(&profile.Password, validation.Required, validation.Length(8, 50), is.Alphanumeric),
	)
}

func SignUp(c *gin.Context) {
	profile := api.SignUp{}
	if err := c.BindJSON(&profile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := ValidProfile(profile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(profile.Password), 7)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user := model.User{
		Username:  profile.Username,
		Email:     profile.Email,
		Password:  string(hash),
		BirthDate: profile.BirthDate,
	}

	if err := user.CreateUser(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, user)
}