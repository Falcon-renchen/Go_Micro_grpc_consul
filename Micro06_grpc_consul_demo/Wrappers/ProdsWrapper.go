package Wrappers

import (
	Services "Go_Micro/Micro06_grpc_consul_demo/Services/protos"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"strconv"
)

type ProdsWrapper struct {
	client.Client
}

func NewLogWrapper(c client.Client) client.Client  {
	return &ProdsWrapper{c}
}

func newProd(id int32, pname string) *Services.ProdsModel {
	return &Services.ProdsModel{
		ProdID:id,
		ProdName:pname,
	}
}

//通用降级方法
func defaultData(rsp interface{})  {
	switch t:=rsp.(type) {
	case *Services.ProdListResponse:
		defaultProds(t)
	case *Services.ProdDetailResponse:
		//t.Data = newProd(10,"降级商品")
	}
}

func defaultProds(rsp interface{})  {
	models := make([]*Services.ProdsModel, 0)
	var i int32
	for i=0; i<5; i++ {
		models = append(models, newProd(100+i,"pname"+strconv.Itoa(100+int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	//result.Data = models
	fmt.Println(result)
}

func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	commandName := req.Service()+"." + req.Endpoint()

	configA := hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 2,
		SleepWindow:            5000,
		ErrorPercentThreshold:  50,
	}
	hystrix.ConfigureCommand(commandName,configA)

	return hystrix.Do(commandName, func() error {
		return this.Client.Call(ctx,req,rsp)

	}, func(err error) error {
		//defaultProds(rsp)
		defaultData(rsp)
		return nil
	})
}




