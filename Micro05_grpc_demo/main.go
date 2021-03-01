package main

import (
	"Go_Micro/Micro05_grpc_demo/ServiceImpl"
	Services "Go_Micro/Micro05_grpc_demo/Services/protos"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
			registry.Addrs("172.16.17.150:8500"),
			)

	prodService := micro.NewService(
			micro.Name("prodservice"),
			micro.Address(":8011"),
			micro.Registry(consulReg),
		)
	prodService.Init()
	//以上是空服务，所以用下面的方法插入服务.
	Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))
	prodService.Run()

}
