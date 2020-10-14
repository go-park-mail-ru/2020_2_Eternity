package jwthelper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
)

type Claims struct {
	Id int `json:"id"`
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
