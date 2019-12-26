package logger

import (
	"io"
	"log"
	"os"
	"time"
)

func Info(info interface{}) {
	//创建日志文件
	f, err := os.OpenFile("golang_"+time.Now().Format("20060102")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	writers := []io.Writer{
		f,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(info)
}