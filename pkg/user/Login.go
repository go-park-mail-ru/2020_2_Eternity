package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	form := api.Login{}

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user := User{
		Username: form.Username,
	}

	if !user.GetUserByName() {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "invalid username"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "bad password"})
		return
	}

	ss, err := jwthelper.CreateJwtToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Error{Error: "cannot create token"})
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
