package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"net/http"
)

func Follow(c *gin.Context) {
	followingUser, followedUser, status, err := PrepareFollow(c)
	if err != nil {
		c.AbortWithStatusJSON(status, err)
	}
	if err := followingUser.Follow(followedUser.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{Error: "Already Followed"})
		return
	}
	c.Status(http.StatusOK)
}

func Unfollow(c *gin.Context) {
	followingUser, followedUser, status, err := PrepareFollow(c)
	if err != nil {
		c.AbortWithStatusJSON(status, err)
	}
	if err := followingUser.UnFollow(followedUser.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{Error: "Already Followed"})
		return
	}
	c.Status(http.StatusOK)
}

func PrepareFollow(c *gin.Context) (*User, *User, int, *Error) {
	followed := api.Follow{}

	if err := c.BindJSON(&followed); err != nil {
		return nil, nil, http.StatusBadRequest, &Error{"bad json"}
	}

	claimsId, ok := GetClaims(c)

	if !ok {
		return nil, nil, http.StatusUnauthorized, &Error{Error: "invalid token"}
	}

	followedUser := User{
		Username: followed.Username,
	}

	if !followedUser.GetUserByName() {
		return nil, nil, http.StatusBadRequest, &Error{Error: "User not found"}
	}

	followingUser := User{
		ID: claimsId,
	}
	return &followingUser, &followedUser, http.StatusOK, nil
}
