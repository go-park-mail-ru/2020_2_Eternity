package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/model"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func SetAvatar(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"Form error"})
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

	filename := filepath.Base(file.Filename)

	root, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"server env error"})
		return
	}

	// TODO: Сгенерировать уникальное имя файла

	path := root + "/static/avatar/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Upload error"})
		return
	}

	if err := user.UpdateAvatar(filename); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, "success")
}

func GetAvatar(c *gin.Context) {
	claims, ok := c.Get("info")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Error{"invalid token"})
		return
	}
	user := model.User{
		ID:       claims.(jwthelper.Claims).Id,
		Username: claims.(jwthelper.Claims).Username,
	}
	if err := user.GetAvatar(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})
		return
	}

	root, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"server env error"})
		return
	}
	path := root + "/static/avatar/" + user.Avatar

	data, err := ioutil.ReadFile(path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Error reading file"})
		return
	}
	c.Header("Content-type", "image/jpeg")
	_, err = c.Writer.Write(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Error write to response"})
		return
	}
	c.Status(http.StatusOK)
}
