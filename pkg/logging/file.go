package logging

import (
	"../setting"
	"fmt"
	"log"
	"os"
	"time"
)

func getLogFullPath() string {
	suffixPath := fmt.Sprintf("%s%s.%s", setting.LogSaveName, time.Now().Format(setting.LogTimeFormat), setting.LogFileExt)
	return fmt.Sprintf("%s%s", setting.LogSavePath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)

	dir, _ := os.Getwd()

	switch {
	case os.IsNotExist(err):
		os.MkdirAll(dir+setting.LogSavePath, os.ModePerm)
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to open file : %v", err)
	}

	return handle

}
