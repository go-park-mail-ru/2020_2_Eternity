package handler

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

// Minimum eight characters, at least one uppercase letter, one lowercase letter and one number:

func ValidProfile(profile model.SignUpUser) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Email, validation.Required, is.EmailFormat),
		validation.Field(&profile.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric),
		validation.Field(&profile.Password, validation.Required, validation.Length(8, 50), is.Alphanumeric),
	)
}

func (h *Handler) SignUp(c *gin.Context) {
	profile := model.SignUpUser{}
	if err := c.BindJSON(&profile); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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
		ID: 0,
		Username: profile.Username,
		Email: profile.Email,
		Password: string(hash),
		Age: profile.Age,
	}

	if err := h.users.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, user)
}