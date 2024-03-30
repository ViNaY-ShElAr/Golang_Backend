package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Setup(level string) {
	// Set the log level based on the configuration
	switch level {
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	fmt.Println("Logger Initialized Sucessfully")
}

func Error(message string, err error) {
	Log.WithFields(logrus.Fields{
		"error": err,
	}).Error(message)
}
