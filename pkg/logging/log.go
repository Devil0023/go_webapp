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

//Setup 注册log
func Setup() {

	filePath := GetLogFileFullPath()

	F = OpenLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)

}

//SetPrefix 设置文件名前缀
func SetPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

//Debug debug级
func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)
}

//Info info级
func Info(v ...interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}

//Warn warn级
func Warn(v ...interface{}) {
	SetPrefix(WARN)
	logger.Println(v)
}

//Error error级
func Error(v ...interface{}) {
	SetPrefix(ERROR)
	logger.Fatalln(v)
}

//Fatal fatal级
func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	logger.Println(v)
}
