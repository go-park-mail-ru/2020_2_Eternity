package handlers

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"net/http"
	"time"
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

	if err := user.UpdateUser(&profile); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Error{err.Error()})
		return
	}

	ss, err := jwthelper.CreateJwtToken(user.ID, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"cannot create token"})
		return
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    ss,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, user)
}
