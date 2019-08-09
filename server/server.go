package server

import (
	"fmt"
	"github.com/fvbock/endless"
	"go_webapp/pkg/setting"
	"go_webapp/routers"
	"log"
	"syscall"
)

//Run 运行Server
func Run() {

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server Error: $v", err)
	}
}
