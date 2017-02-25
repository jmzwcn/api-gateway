package log

import "log"

//TODO
func Debug(v ...interface{}) {
	log.Println(v)
}

func Info(v ...interface{}) {
	log.Println(v)
}

func Warn(v ...interface{}) {
	log.Println(v)
}

func Error(v ...interface{}) {
	log.Println(v)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}
