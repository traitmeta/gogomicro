package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic"
	"github.com/songxuexian/gogomicro/basic/common"
	"github.com/songxuexian/gogomicro/basic/config"
	"github.com/songxuexian/gogomicro/orders-srv/handler"
	"github.com/songxuexian/gogomicro/orders-srv/model"
	"github.com/songxuexian/gogomicro/orders-srv/subscriber"
	"time"

	proto "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
)

func main() {
	basic.Init()
	micReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.orders"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) {
			model.Init()
			handler.Init()
			subscriber.Init()
		}))

	// Register Handler
	err := proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
	if err != nil {
		log.Fatal(err)
	}

	// Register Struct as Subscriber
	err = micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
