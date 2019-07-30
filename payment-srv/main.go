package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic"
	"github.com/songxuexian/gogomicro/basic/config"
	"github.com/songxuexian/gogomicro/payment-srv/handler"
	"github.com/songxuexian/gogomicro/payment-srv/model"
	"time"

	payment "github.com/songxuexian/gogomicro/payment-srv/proto/payment"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.payment"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	_ = payment.RegisterPaymentHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
