package main

import (
	model "Go_Micro/Micro06_grpc_consul_demo/model/protos"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"net/http"
)




func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("172.16.17.144:8500"),
		)
/*
	service, err := consulReg.GetService("productservice")
	if err != nil {
		panic(err)
	}
	//fmt.Println(service)
	next := selector.Random(service)
	node, err := next()
	//fmt.Println(node.Id,node.Address,node.Metadata)
	callApi, err := CallApi(node.Address, "/v1/prods", "GET")
	fmt.Println(callApi)
*/

}

//原始调用
func CallApi(addr string, path string, method string) (string, error) {
	req,_ := http.NewRequest(method,"http://"+addr+path,nil)
	client := http.DefaultClient
	res, _ := client.Do(req)
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}


func CallApi2(s selector.Selector)  {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
		)
	req := myClient.NewRequest("productservice", "/v1/prods", model.ProdsRequest{Size:3})
	var rsp model.ProdListResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetData())
}