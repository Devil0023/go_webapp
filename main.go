package main

import (
	"go_webapp/cron"
	"go_webapp/models"
	"go_webapp/pkg/gredis"
	"go_webapp/pkg/logging"
	"go_webapp/pkg/setting"
	"go_webapp/server"
	"runtime"
	"sync"
)

//main main函数
func main() {

	// 初始化模块
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logging.Setup()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	//crontab
	go func() {
		defer wg.Done()
		cron.Run()
	}()

	//server
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()

}
