package main

import (
	"category/common"
	"category/config"
	"category/handler"
	pb "category/proto"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"
	_"category/model"
	"github.com/micro/go-micro/v2"

)

func main() {
	//配置中心
	consulconfig,err:=config.GetConsulConfig("127.0.0.1",8500,"/lizhengda/project/config")
	if err!=nil{
		log.Error(err)
	}
	//注册中心
	registry:=consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs=[]string{
				"127.0.0.1:8500",
			}

		},
	)

	// Create service
	srv := micro.NewService(
		micro.Name("micro.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8888"),
		//设置配置中心
		micro.Config(consulconfig),
		//设置注册中心
		micro.Registry(registry),
	)

	// Register handler
	pb.RegisterCategoryHandler(srv.Server(), new(handler.Category))
	common.GetMysqlConfigFromConsul(consulconfig,"mysql")

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
