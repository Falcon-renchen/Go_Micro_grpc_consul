package handler

import (
	sum "Go_Micro/Micro02/proto/sum"
)
import "context"
import "Go_Micro/Micro02/sum-srv/server"

type handler struct {

}

func (h handler) GetSum(ctx context.Context,req *sum.SumRequest,rsp *sum.SumResponse) error {
	inputs := make([]int64, 0)

	var i int64 = 0
	for ; i<=req.Intput; i++ {
		inputs = append(inputs, i)
	}

	rsp.Output = server.GetSum(inputs...)
	return nil
}

func Handler() sum.SumHandler {
	return handler{}
}



