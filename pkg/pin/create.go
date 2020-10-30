package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FormCreatePin struct {
	CreatePinApi *api.CreatePin        `form:"data" binding:"required"`
	Avatar       *multipart.FileHeader `form:"img" binding:"required"`
}

func CreatePin(c *gin.Context) {
	claimsId, ok := auth.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Println("[CreatePin]: Can't get claims")
		return
	}

	formPin := FormCreatePin{}
	if err := c.Bind(&formPin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("[CreatePin]-[Bind] :" + err.Error())
		return
	}

	fileName, err := utils.RandomUuid()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("[CreatePin]-[RandomUuid] :" + err.Error())
		return
	}

	if err := os.MkdirAll(config.Conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println("[CreatePin]-[MkAllDir] :" + err.Error())
		return
	}

	if err := c.SaveUploadedFile(formPin.Avatar, config.Conf.Web.Static.DirImg+"/"+fileName); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println("[CreatePin]-[SaveUploadedFile] :" + err.Error())
		return
	}

	pin := Pin{
		Title:       formPin.CreatePinApi.Title,
		Content:     formPin.CreatePinApi.Content,
		PictureName: fileName,
		UserId:      claimsId,
	}

	if err := pin.CreatePin(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println("[CreatePin]-[CreatePin] :" + err.Error())
		return
	}

	log.Printf("[CreatePin]: pin{%v %v %v %v %v}", pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId)

	c.JSON(http.StatusOK, api.GetPin{
		Id:      pin.Id,
		Title:   pin.Title,
		Content: pin.Content,
		ImgLink: filepath.Join(config.Conf.Web.Static.UrlImg, pin.PictureName), // TODO full path
		UserId:  pin.UserId,
	})
}
