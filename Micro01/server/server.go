package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
)
import proto "Go_Micro/Micro01/proto"

type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	fmt.Println("Hello, Client connecting...")
	rsp.Greeting = "您好," + req.Name
	return nil
}

func main() {
	//创建服务
	service := micro.NewService(
		micro.Name("greeter.service"),
		micro.Version("latest"),
		)
	service.Init()

	//注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	//启动服务
	if err := service.Run(); err!=nil {
		fmt.Println(err)
	}
}
