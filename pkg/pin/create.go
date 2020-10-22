package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type FormCreatePin struct {
	CreatePinApi *api.CreatePin        `form:"data" binding:"required"`
	Avatar       *multipart.FileHeader `form:"img" binding:"required"`
}

func CreatePin(c *gin.Context) {
	claimsId, ok := user.GetClaims(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		config.Lg("pin", "CreatePin").Error("Can't get claims")
		return
	}

	formPin := FormCreatePin{}
	if err := c.Bind(&formPin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin", "CreatePin").Error("Bind: ", err.Error())
		return
	}

	fileName, err := utils.RandomUuid()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		config.Lg("pin", "CreatePin").Error("RandomUuid: ", err.Error())
		return
	}

	if err := os.MkdirAll(config.Conf.Web.Static.DirImg, 0777|os.ModeDir); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		config.Lg("pin", "CreatePin").Error("MkAllDir: ", err.Error())
		return
	}

	if err := c.SaveUploadedFile(formPin.Avatar, config.Conf.Web.Static.DirImg+"/"+fileName); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		config.Lg("pin", "CreatePin").Error("SaveFile: ", err.Error())
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
		config.Lg("pin", "CreatePin").Error("pin.CreatePin: ", err.Error())
		return
	}

	config.Lg("pin", "CreatePin").
		Infof("Created pin {%v %v %v %v %v}", pin.Id, pin.Title, pin.Content, pin.PictureName, pin.UserId)

	imgUrl := url.URL{
		Scheme: config.Conf.Web.Server.Protocol,
		Host:   config.Conf.Web.Server.Host,
		Path:   filepath.Join(config.Conf.Web.Static.UrlImg, pin.PictureName),
	}

	c.JSON(http.StatusOK, api.GetPin{
		Id:      pin.Id,
		Title:   pin.Title,
		Content: pin.Content,
		ImgLink: imgUrl.String(),
		UserId:  pin.UserId,
	})
}
