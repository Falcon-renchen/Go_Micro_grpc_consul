package main

import (
	"Go_Micro/Micro02/prime-srv/handler"
	"Go_Micro/Micro02/proto/prime"
	"github.com/go-log/log"
	"github.com/micro/go-micro"
)

func main() {
	//创建服务
	srv := micro.NewService(
		micro.Name("prime"),
		)
	//初始化
	srv.Init(
		micro.BeforeStart(func() error {
			log.Log("srv 启动前的日志")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log("srv 启动后的日志")
			return nil
		}),
		)

	//	挂载接口
	_ = prime.RegisterPrimeHandler(srv.Server(), handler.Handler())

	if err := srv.Run(); err!=nil {
		panic(err)
	}
}
