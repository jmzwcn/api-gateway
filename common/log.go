package log

import (
	log "github.com/sirupsen/logrus"
)

func Debug(v ...interface{}) {
	log.Debug(v)
}

func Info(v ...interface{}) {
	log.Info(v)
}

func Warn(v ...interface{}) {
	log.Warn(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}

func SetLevel(level string) {
	if level, err := log.ParseLevel(level); err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}
