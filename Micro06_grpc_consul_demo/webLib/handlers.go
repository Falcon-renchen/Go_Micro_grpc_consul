package webLib

import (
	"github.com/gin-gonic/gin"
	"strconv"
)
import "context"
import Services "Go_Micro/Micro06_grpc_consul_demo/Services/protos"

func newProd(id int32, pname string) *Services.ProdsModel {
	return &Services.ProdsModel{
		ProdID:id,
		ProdName:pname,
	}
}

func defaultProds() (*Services.ProdListResponse, error) {
	models := make([]*Services.ProdsModel, 0)
	var i int32
	for i=0; i<5; i++ {
		models = append(models, newProd(100+i,"pname"+strconv.Itoa(100+int(i))))
	}
	res := &Services.ProdListResponse{}
	//res.Data = models
	return res, nil
}

func PanicIfError(err error)  {
	if err != nil {
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context)  {
	//绑定
	var prodReq Services.ProdsRequest
	PanicIfError(ginCtx.BindUri(&prodReq))
	prodService := ginCtx.Keys["prodservie"].(Services.ProdService)
	resp, _ := prodService.GetProdDetail(context.Background(),&prodReq)
	ginCtx.JSON(200,gin.H{
		"data" : resp.Data,
	})
}

//gin的方法部分
func GetProdsList(ginCtx *gin.Context)  {
	prodService := ginCtx.Keys["prodservie"].(Services.ProdService)
		var prodReq Services.ProdsRequest
		err := ginCtx.Bind(&prodReq)

	if err!=nil {
		ginCtx.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {
		prodRes, _ := prodService.GetProdsList(context.Background(), &prodReq)

		//if err != nil {
		//	ginCtx.JSON(500, gin.H{
		//		"status": err.Error(),
		//	})
		//} else {
			ginCtx.JSON(200, gin.H{
				"data": prodRes.Data,
			})
		//}
		/*
				//临时代码而已
				//熔断代码改造
				//第一步 配置config
				configA := hystrix.CommandConfig{
					Timeout:                1000,
					MaxConcurrentRequests:  0,
					RequestVolumeThreshold: 0,
					SleepWindow:            0,
					ErrorPercentThreshold:  0,
				}
				//第二步 配置command
				hystrix.ConfigureCommand("getprods", configA)
				var prodRes *Services.ProdListResponse
				//第三步 执行使用do方法
				err = hystrix.Do("getprods", func() error {
					prodRes, err = prodService.GetProdsList(context.Background(), &prodReq)
					return err
				}, func(err error) error {
					prodRes, err = defaultProds()
					return err
				})

				if err != nil {
					ginCtx.JSON(500, gin.H{
						"status": err.Error(),
					})
				} else {
					ginCtx.JSON(200, gin.H{
						"data": prodRes.Data,
					})
				}

			}
		*/
	}

}