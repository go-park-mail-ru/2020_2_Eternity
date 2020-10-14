package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func SetAvatar(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"Form error"})
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

	root, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"server env error"})
		return
	}

	filename, err := utils.RandomUuid()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Random cant generate UUID"})
	}

	path := root + config.Conf.Web.Static.DirAvt + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Upload error"})
		return
	}

	if err := user.UpdateAvatar(utils.GenerateUrlAvatar(filename)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, "success")
}

func GetAvatar(c *gin.Context) {
	root, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"server env error"})
		return
	}
	filename := c.Param("file")

	path := root + config.Conf.Web.Static.DirAvt + filename

	data, err := ioutil.ReadFile(path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"Error filename"})
		return
	}
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	_, err = c.Writer.Write(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Error{"Error write to response"})
		return
	}
	c.Status(http.StatusOK)
}
