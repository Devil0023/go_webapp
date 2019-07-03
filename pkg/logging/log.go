package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func init() {

	filePath := getLogFileFullPath()

	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)

}

func SetPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

//func Debug(v ...interface{}){
//	SetPrefix(DEBUG)
//	logger.Println(v)
//}
