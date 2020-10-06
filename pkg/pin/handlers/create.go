package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin/model"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func randomUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	} else {
		return u.String(), err
	}
}

func generateRelPath(filename string) string {
	fn := []rune(filename)

	const Depth int = 5  // to config
	const DirNameLen int = 1 // to config

	var realDepth int
	if Depth * DirNameLen > len(fn) {
		realDepth = len(fn) / DirNameLen
	} else {
		realDepth = Depth / DirNameLen
	}

	if realDepth <= 0 {
		realDepth = 1
	}

	dirs := []string{}
	for i := 0; i < realDepth; i += DirNameLen {
		dirs = append(dirs, string(fn[i:i+DirNameLen]))
	}

	return strings.Join(dirs, "/") + "/" + filename
}



func prepareFileStorage() (filePath string, err error) {
	u, err := randomUuid()
	if err != nil {
		return "", err
	}

	relPath := generateRelPath(u)
	err = os.MkdirAll(filepath.Dir(relPath), os.ModeDir)
	if err != nil {
		return "", err
	}

	const StaticDir = ""  // config

	return StaticDir + "/" + relPath, nil
}


func CreatePin(c *gin.Context) {
	claims, ok := c.Get("info")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("can't get key"))
		return
	}

	requester, ok := claims.(jwthelper.Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.New("can't lead claims"))
		return
	}


	file, err := c.FormFile("file") // ???
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	//filename := filepath.Base(file.Filename)
	filePath, err := prepareFileStorage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	pin := model.Pin{}
	if err := c.BindJSON(&pin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	pin.UserId = requester.Id
	if err := pin.CreatePin(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "{}") // return ID?
}