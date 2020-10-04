package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"net/http"
	"time"
)

const secretkey string = "SECRET"
const cookiename string = "eternity"

type Error struct {
	Error string `json:"error"`
}

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	form := api.Login{}

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user := model.User{
		Username: form.Username,
		Password: form.Password,
	}

	id, exists := user.CheckUser()
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"invalid login or password"})
		return
	}

	ss, err := CreateJwtToken(id, form.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"cannot create token"})
		return
	}

	cookie := http.Cookie{
		Name:     cookiename,
		Value:    ss,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, id)
}

func CreateJwtToken(id int, username string) (string, error) {
	SecretKey := []byte(secretkey)
	claims := Claims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return ss, err
}
