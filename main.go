package main

import (
	"./pkg/setting"
	"./routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", setting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server Error: $v", err)
	}
}
