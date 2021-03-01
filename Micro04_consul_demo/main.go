package main

import (
	"Go_Micro/Micro04_consul_demo/ServiceImpl"
	Services "Go_Micro/Micro04_consul_demo/Services/protos"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("172.16.17.144:8500"))

	myService := micro.NewService(
		micro.Name("test.Go_Micro06"),
		micro.Address(":8001"),
		micro.Registry(consulReg),
		)

	//注册服务
	Services.RegisterTestServiceHandler(myService.Server(),new(ServiceImpl.TestService))
	//运行服务
	myService.Run()
}
