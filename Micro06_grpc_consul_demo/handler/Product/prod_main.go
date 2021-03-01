package main

import (
	"Go_Micro/Micro06_grpc_consul_demo/Services/protos"
	"Go_Micro/Micro06_grpc_consul_demo/Wrappers"
	"Go_Micro/Micro06_grpc_consul_demo/webLib"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

type logWrapper struct {
	client.Client
}

//装饰
func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opt ...client.CallOption) error {
	fmt.Println("调用接口")
	return this.Client.Call(ctx,req,rsp)//调用父类方法
}

func NewLogWrapper(c client.Client) client.Client  {
	return &logWrapper{c}
}




func main() {

	//注册consul
	consulReg := consul.NewRegistry(
		registry.Addrs("172.16.17.144:8500"),
	)

	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper),
		micro.WrapClient(Wrappers.NewLogWrapper),
	)
	prodService := Services.NewProdService("prodservice",myService.Client())



	//1.先做一个web api
	httpserver := web.NewService(
		web.Name("httpservice"),
		web.Address(":8001"),
		web.Handler(webLib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	httpserver.Init()
	httpserver.Run()

}
