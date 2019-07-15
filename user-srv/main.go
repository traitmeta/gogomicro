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
	"github.com/songxuexian/gogomicro/user-srv/handler"
	"github.com/songxuexian/gogomicro/user-srv/models"
	user "github.com/songxuexian/gogomicro/user-srv/proto/user"
	"time"
)

func main() {
	basic.Init()

	micReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			models.Init()
			handler.Init()
		}),
	)

	// Register Handler
	err := user.RegisterUserHandler(service.Server(), new(handler.Service))
	if err != nil {
		log.Fatal(err)
	}

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("sxx.micro.book.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("sxx.micro.book.srv.user", service.Server(), subscriber.Handler)

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
