package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("172.16.17.144:8500"),
		)

	ginRouter := gin.Default()
	ginRouter.Handle("GET","/index", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(200,gin.H{
			"data":data,

		})
	})

	

	//1.先做一个web api
	server := web.NewService(
		web.Name("main"),
		web.Address(":8000"),
		web.Handler(ginRouter),
		web.Registry(consulReg),

	)

	server.Run()
}
