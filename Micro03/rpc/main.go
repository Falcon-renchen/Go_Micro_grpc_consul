package main

import (
	"Go_Micro/Micro03/proto/learning"
	greeterH "Go_Micro/Micro03/rpc/handler/greeter"
	learningH "Go_Micro/Micro03/rpc/handler/learning"
	"github.com/micro/go-micro"
	"log"
)

///以下启动rpc服务
func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.learning"),
		)

	//注册Greeter接口
	err := learning.RegisterGreeterHandler(service.Server(),new(greeterH.Handler))
	if err != nil {
		log.Fatal(err)
	}


	//注册learning接口
	err = learning.RegisterLearningHandler(service.Server(),new(learningH.Handler))
	if err != nil {
		log.Fatal(err)
	}

	service.Init()

	log.Fatal(service.Run())

}
///以上启动rpc服务
