package handlers

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ValidPasswords(pswds api.UpdatePassword) error {
	return validation.ValidateStruct(&pswds,
		validation.Field(&pswds.OldPassword, validation.Required, validation.Length(8, 50), is.Alphanumeric),
		validation.Field(&pswds.NewPassword, validation.Required, validation.Length(8, 50), is.Alphanumeric),
	)
}

func UpdatePassword(c *gin.Context) {
	psswds := api.UpdatePassword{}

	if err := c.BindJSON(&psswds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := ValidPasswords(psswds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	claims, ok := c.Get("info")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"invalid token"})
		return
	}

	user := model.User{
		ID:       claims.(jwthelper.Claims).Id,
		Username: claims.(jwthelper.Claims).Username,
	}

	exists := user.GetUser()

	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"user doesnt exist"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(psswds.OldPassword)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"bad password"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(psswds.NewPassword), config.Conf.Token.Value)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := user.UpdatePassword(string(hash)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
