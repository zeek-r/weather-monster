package logger

import (
	"io"

	logger "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

func InitLogger(output io.Writer) {
	// Set JSON as log formt
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(output)
}

func Debug(message string, args ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if args != nil {
		fields = args[0]
	}
	logger.WithFields(fields).Debug(message)
}

func Info(message string, args ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if args != nil {
		fields = args[0]
	}
	logger.WithFields(fields).Info(message)
}

func Fatal(err error, message string, args ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if args != nil {
		fields = args[0]
	}
	logger.WithFields(fields).WithError(err).Fatal(message)
}

func Error(err error, message string, args ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if args != nil {
		fields = args[0]
	}
	logger.WithFields(fields).WithError(err).Error(message)
}
