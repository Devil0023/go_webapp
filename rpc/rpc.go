package rpc

import (
	"fmt"
	"go_webapp/pkg/logging"
	"go_webapp/pkg/setting"
	"go_webapp/rpc/proto"
	"go_webapp/rpc/services"
	"google.golang.org/grpc"
	"net"
)

func Run() {
	listen, err := net.Listen("tcp", setting.RpcSetting.Listen)

	if err != nil {
		logging.Info("failed to listen: ", err)
	}

	server := grpc.NewServer()

	// 注册服务
	registerRpcServices(server)

	fmt.Println("Start to serve")

	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}
}

func registerRpcServices(server *grpc.Server) {

	proto.RegisterHelloServer(server, services.HelloService)

}
