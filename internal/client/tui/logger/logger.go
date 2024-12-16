package logger

import (
	"log"
	"os"
)

func Get() *log.Logger {
	f, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println("can't set logger")
		return nil
	}
	return log.New(f, "", 0)
}
