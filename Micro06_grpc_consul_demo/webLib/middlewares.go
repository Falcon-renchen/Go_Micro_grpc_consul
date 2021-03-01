package webLib

import (
	Services "Go_Micro/Micro06_grpc_consul_demo/Services/protos"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservice"] = prodService //赋值
		context.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r:=recover(); r!=nil {
				context.JSON(500,gin.H{
					"status":fmt.Sprintf("%s",r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}