package config

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"strings"
)

type Logger struct {
	logFile *os.File
}

func (l *Logger) Init() {
	if !Conf.Logger.StdoutLog {
		file, err := os.OpenFile(Conf.Logger.CommonFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
			l.logFile = file
		} else {
			Lg("config", "Logger.Init").Error("Failed to log to file, using default stderr")
		}
	}

	log.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	switch strings.ToLower(Conf.Logger.CommonLevel) {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	Lg("config", "Logger.Init").
		Info("Created logger " + strings.ToUpper(log.GetLevel().String()) + " level")
}

func (l *Logger) Cleanup() error {
	return l.logFile.Close()
}

func Lg(packageName, functionName string) *log.Entry {
	return log.WithFields(log.Fields{
		"package":  packageName,
		"function": functionName,
	})
}
