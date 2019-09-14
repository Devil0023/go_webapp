package services

import (
	"context"
	"go_webapp/rpc/proto"
)

type helloService struct {
}

var HelloService = helloService{}

func (s helloService) SayHello(ctx context.Context, in *proto.HelloRequest) (out *proto.HelloReply, err error) {

	name := in.Name

	resp := new(proto.HelloReply)

	resp.Message = "Hello: " + name

	return resp, nil
}
