package jwthelper

import (
	b64 "encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"time"
)

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type CsrfClaims struct {
	Id      int
	Val     string
	Expires time.Time
	jwt.StandardClaims
}

func CreateJwtToken(id int) (string, error) {
	SecretKey := []byte(config.Conf.Token.SecretName)
	claims := Claims{
		Id: id,
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

func ParseToken(cookie string, claims *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad token signing method")
		}
		return []byte(config.Conf.Token.SecretName), nil
	})
}

func CreateCsrfToken(id int, val string, expires time.Time) (string, error) {
	SecretKey := []byte(config.Conf.Token.SecretName)
	claims := CsrfClaims{
		Id:      id,
		Val:     val,
		Expires: expires,
		StandardClaims: jwt.StandardClaims{
			Issuer: "server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString([]byte(ss)), err
}

func ParseCsrfToken(value string, claims *CsrfClaims) (*jwt.Token, error) {
	bvalue, err := b64.URLEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	return jwt.ParseWithClaims(string(bvalue), claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad token signing method")
		}
		return []byte(config.Conf.Token.SecretName), nil
	})
}

func GetClaims(c *gin.Context) (int, bool) {
	claims, ok := c.Get("info")
	if !ok {
		return -1, false
	}

	claimsId, ok := claims.(int)
	if !ok {
		return -1, false
	}
	return claimsId, true
}
