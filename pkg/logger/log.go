package logger

import (
	"log"
	"os"
)

func InitLog(path string) {
	fileName := "log.log"
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	file, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
