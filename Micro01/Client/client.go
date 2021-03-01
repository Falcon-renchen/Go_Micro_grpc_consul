package main

import (
	proto "Go_Micro/Micro01/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)
func main() {
	//定义服务
	service := micro.NewService(
		micro.Name("greeter.client"),
		)
	service.Init()

	//创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	//调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro wyp"})
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印相应结果
	fmt.Println(rsp.Greeting)
}
