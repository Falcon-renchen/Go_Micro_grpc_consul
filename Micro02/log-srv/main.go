package main

import (
	proto "Go_Micro/Micro02/proto/log"
	"context"
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

//监听broker异步消息

type Sub struct {

}

func (s *Sub) Process(ctx context.Context, evt *proto.LogEvt) error {
	//业务逻辑
	log.Infof("收到日志: %v", evt.Msg)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.srv.log"),
		micro.Version("latest"),
		)
	service.Init()

	micro.RegisterSubscriber("go.micro.learning.topic.log",service.Server(),&Sub{})

	service.Run()
}
