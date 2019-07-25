package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/songxuexian/gogomicro/orders-srv/handler"
	"github.com/songxuexian/gogomicro/orders-srv/subscriber"

	order "github.com/songxuexian/gogomicro/orders-srv/proto/order"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.order", service.Server(), new(subscriber.Order))

	// Register Function as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.order", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
