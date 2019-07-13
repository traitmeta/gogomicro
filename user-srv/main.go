package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/user-srv/handler"
	user "github.com/songxuexian/gogomicro/user-srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := user.RegisterUserHandler(service.Server(), new(handler.User))
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
