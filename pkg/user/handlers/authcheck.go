package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(cookiename)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"bad cookie"})
			return
		}
		claims := Claims{}
		token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("bad token signing method")
			}
			return []byte(secretkey), nil
		})
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"bad token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"fake token"})
			return
		}
		c.Set("info", claims)
		c.Next()
	}
}
