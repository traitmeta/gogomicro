package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/auth/handler"
	"github.com/songxuexian/gogomicro/auth/model"
	"github.com/songxuexian/gogomicro/basic"
	"github.com/songxuexian/gogomicro/basic/config"
	"time"

	auth "github.com/songxuexian/gogomicro/auth/proto/auth"
)

func main() {
	basic.Init()
	micReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// Register Handler
	_ = auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

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
