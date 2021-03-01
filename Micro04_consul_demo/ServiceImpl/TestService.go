package ServiceImpl

import (
	Services "Go_Micro/Micro04_consul_demo/Services/protos"
	"strconv"
)
import "context"

type TestService struct {

}

func (t *TestService) Call(ctx context.Context, req *Services.TestRequest, rsp *Services.TestResponse)  error {
	rsp.Data = "text" + strconv.Itoa(int(req.Id))
	return nil
}

