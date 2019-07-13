package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"./user-srv/handler"
	"./user-srv/subscriber"

	user "./user-srv/proto/user"
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
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
