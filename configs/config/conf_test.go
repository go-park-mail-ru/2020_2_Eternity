package config

import (
	"os"
	"testing"
)

var (
	logLevel = []string {
		"trace",
		"debug",
		"info",
		"warn",
		"error",
		"fatal",
		"panic",
	}

)

var _ = func() bool {
	testing.Init()
	Conf = NewConfigTst()
	return true
}()

func TestLoggerNew(t *testing.T) {
	Conf.Logger.StdoutLog = true
	l := Logger{
		os.Stdout,
	}

	lvlWas := Conf.Logger.CommonLevel
	for _, lvl := range logLevel {
		Conf.Logger.CommonLevel = lvl
		l.Init()
	}

	Conf.Logger.CommonLevel = lvlWas
}

func TestConfSet(t *testing.T) {
	setDefaultWeb()
	setDefaultDb()
	setDefaultLog()
}

func TestConf(t *testing.T) {
	setDefaultWeb()
	setDefaultDb()
	setDefaultLog()
	splitPath("path")
}
