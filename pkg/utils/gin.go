package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"strconv"
)

func GetIntParam(c *gin.Context, param string) (int, error) {
	value, ok := c.Params.Get(param)
	if !ok {
		config.Lg("comment", "GetIntParam").Info("Param " + param + " not found in url")
		return 0, errors.New("Param " + param + " not found in url")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		config.Lg("comment", "GetIntParam").Info(err.Error())
		return 0, err
	}

	return intValue, nil
}
