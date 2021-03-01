package main

import (
	logProto "Go_Micro/Micro02/proto/log"
	"Go_Micro/Micro02/proto/sum"
	"Go_Micro/Micro02/sum-srv/handler"
	"context"
	"encoding/json"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		micro.WrapHandler(
			//并行请求
			rateLimiter(5),
			),
			micro.Flags(&cli.StringFlag{
				Name:        "learning_go",
				Usage:       "help一下，你就知道",
				EnvVar:      "",
				Hidden:      false,
				Value:       "",
				Destination: nil,
			}),
		)
	//初始化
	srv.Init(
		micro.WrapHandler(
			reqLogger(srv.Client()),
			),
		micro.BeforeStart(func() error {
			log.Error("srv 启动前的日志")
			return nil
			}),
			micro.AfterStart(func() error {
				log.Error("srv 启动后的日志")
				return nil
			}),
		)

	//挂载接口
	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())

	//启动
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
//接收  transport
//发送  broker

func rateLimiter(rate int) server.HandlerWrapper {
	tokens := make(chan bool, rate)
	for i:=0; i<rate; i++ {
		tokens <- true
	}
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			token := <- tokens
			defer func() {
				tokens <- token
			}()
			return handlerFunc(ctx,req,rsp)
		}
	}
}


//实现日志的wrapper
//通过Broker异步消息日志推送到日志服务
func reqLogger(cli client.Client) server.HandlerWrapper {
	//推送器
	pub := micro.NewPublisher("go.micro.learning.topic.log",cli)
	//初始化动作
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		//中间动作
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Info("发送日志")
			evt := logProto.LogEvt {
				Msg : "Hello",
			}
			body, _ := json.Marshal(evt)
			pub.Publish(ctx, &broker.Message{
				Header: map[string]string{
					"serviceName":"sum",
				},
				Body:   body,
			})
			return handlerFunc(ctx, req, rsp)
		}
	}
}