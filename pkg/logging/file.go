package logging

import (
	"fmt"
	"go_webapp/pkg/setting"
	"log"
	"os"
	"time"
)

func getLogFileFullPath() string {

	suffixPath := fmt.Sprintf(
		"%s%s.%s",
		setting.LogSetting.LogSaveName,
		time.Now().Format(setting.LogSetting.LogTimeFormat),
		setting.LogSetting.LogFileExt)

	return fmt.Sprintf("%s%s", setting.LogSetting.LogSavePath, suffixPath)

}

func openLogFile(filePath string) *os.File {

	dir, _ := os.Getwd()

	_, err := os.Stat(filePath)

	switch {

	case os.IsNotExist(err):
		os.MkdirAll(dir+setting.LogSetting.LogSavePath, os.ModePerm)
		break

	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
		break

	}

	handle, err := os.OpenFile(dir+filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to open file : %v", err)
	}

	return handle

}
