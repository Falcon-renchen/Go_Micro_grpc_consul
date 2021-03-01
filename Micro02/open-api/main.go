package main

import (
	"Go_Micro/Micro02/proto/prime"
	"Go_Micro/Micro02/proto/sum"
	"context"
	"encoding/json"
 	"github.com/micro/cli"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	proto "Go_Micro/Micro02/proto/api"
	"strconv"
)

var sumClient sum.SumService
var primeClient prime.PrimeService

type Open struct {
	
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.aip.open"),
	)
	service.Init(micro.Action(func(c *cli.Context) {
		sumClient = sum.NewSumService("go.micro.learning.web.sum",service.Client())
		primeClient = prime.NewPrimeService("go.micro.learning.web.prime",service.Client())
		return
	}))

	//暴露接口
	_ = proto.RegisterOpenHandler(service.Server(),new(Open))

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func (open *Open) Fetch(ctx context.Context, req *api.Request, rsp *api.Response) error {
	sumInputStr := req.Get["sum"].Values[0]
	primeInputStr := req.Get["prime"].Values[0]

	sumInput, _ := strconv.ParseInt(sumInputStr, 10, 10)
	primeInput, _ := strconv.ParseInt(primeInputStr, 10, 10)

	sumReq := &sum.SumRequest{Intput:sumInput,}
	primeReq := &prime.PrimeRequest{Input:primeInput}


	//调用客户端
	sumRsp, _ := sumClient.GetSum(context.Background(), sumReq)
	primeRsp, _ := primeClient.GetPrime(context.Background(), primeReq)

	//编码
	ret, _ := json.Marshal(map[string]interface{}{
		"sum" : sumRsp.Output,
		"prime" : primeRsp.Output,
	})

	rsp.Body = string(ret)

	return  nil
}


