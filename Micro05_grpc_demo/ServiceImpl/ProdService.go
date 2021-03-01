package ServiceImpl

import (
	Services "Go_Micro/Micro05_grpc_demo/Services/protos"
	"context"
	"fmt"
	"strconv"
	"time"
)

type ProdService struct {
	
}

func newProd(id int32, pname string) *Services.ProdsModel {
	return &Services.ProdsModel{
		ProdID:id,
		ProdName:pname,
	}
}



func (*ProdService) GetProdsList(ctx context.Context, in *Services.ProdsRequest, res *Services.ProdListResponse) error {
	time.Sleep(time.Second*3)
	models := make([]*Services.ProdsModel, 0)
	var i int32
	for i=0; i<in.Size; i++ {
		models = append(models, newProd(100+i,"prodname"+strconv.Itoa(100+int(i))))
	}
	fmt.Println(models)
	//res.Data = models
	return nil
}

func (*ProdService) GetProdDetail(ctx context.Context, req *Services.ProdsRequest, rsp *Services.ProdDetailResponse) error {
	//rsp.Data = newProd(req.ProdId, "商品详细列表")
	return nil
}

