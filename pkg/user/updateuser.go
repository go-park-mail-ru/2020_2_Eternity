package user

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"net/http"
)

// будет работать, если меняем и email, и username

func ValidUpdate(profile api.UpdateUser) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Email, validation.Required, is.EmailFormat),
		validation.Field(&profile.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric),
	)
}

func UpdateUser(c *gin.Context) {
	profile := api.UpdateUser{}
	if err := c.BindJSON(&profile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := ValidUpdate(profile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	claimsId, ok := GetClaims(c)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"invalid token"})
		return
	}

	user := User{
		ID: claimsId,
	}

	if err := user.UpdateUser(&profile); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Error{err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
