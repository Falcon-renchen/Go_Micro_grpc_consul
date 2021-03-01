package handler

import (
	"Go_Micro/Micro02/prime-srv/server"
	"Go_Micro/Micro02/proto/prime"
	"context"


)

type handler struct {

}

func (h handler) GetPrime(ctx context.Context,req *prime.PrimeRequest,rsp *prime.PrimeResponse) error {
	inputs := make([]int64, 0)

	var i int64 = 0
	for ; i<=req.Input; i++ {
		inputs = append(inputs, i)
	}

	rsp.Output = server.GetPrime(inputs...)
	return nil
}


func Handler() handler {
	return handler{}
}

