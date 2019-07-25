package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic"
	"github.com/songxuexian/gogomicro/basic/config"
	"time"

	"github.com/gogomicro/user-web/handler"
	"github.com/micro/go-micro/web"
)

func main() {
	basic.Init()

	micReg := consul.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("sxx.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8089"),
	)

	// initialise service

	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
