package pin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"log"
	"net/http"
)


func CreatePin(c *gin.Context) {
	claims, ok := user.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, user.Error{"can't get key"})
		return
	}

	file, err := c.FormFile("img") // config
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[FormFile] :" + err.Error()})
		return
	}

	fileName, err := utils.RandomUuid()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[prepareFileStorage]: " + err.Error()})
		return
	}

	if err := c.SaveUploadedFile(file, config.Conf.Web.Static.DirImg+"/"+fileName); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[SaveUploadedFile]: " + err.Error()})
		return
	}

	pinApi := api.CreatePinApi{}
	jsonStr := c.PostForm("data") // config
	log.Print(jsonStr)

	if err := json.Unmarshal([]byte(jsonStr), &pinApi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[Unmarshal]: " + err.Error()})
		return
	}

	pin := Pin{
		Title:   pinApi.Title,
		Content: pinApi.Content,
		ImgLink: config.Conf.Web.Static.UrlImg + "/" + fileName,
		UserId:  claims.Id,
	}

	if err := pin.CreatePin(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, user.Error{"[CreatePin]: " + err.Error()})
		return
	}

	log.Printf("pin{%v %v %v %v %v}", pin.Id, pin.Title, pin.Content, pin.ImgLink, pin.UserId)

	c.JSON(http.StatusOK, "")
}
