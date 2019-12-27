package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

func Info(msg ...interface{}) {
	f, err := os.OpenFile("golang_"+time.Now().Format("20060102")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer closed(f)
	writers := []io.Writer{
		f,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	prefix := "[" + time.Now().Format("2006-01-02 15:04:05.000000") + "]" + " " + "Logger.INFO" + " "
	logger := log.New(fileAndStdoutWriter,prefix,0)
	//logger := log.New(fileAndStdoutWriter,prefix,log.Ldate|log.Ltime|log.Lmicroseconds)
	_, file, line, _ := runtime.Caller(1)
	file_line := fmt.Sprintf("(%v:%v)", path.Base(file), line)
	logger.Println(file_line,msg)
}

func closed(f *os.File)  {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}