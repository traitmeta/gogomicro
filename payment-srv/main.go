package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/songxuexian/gogomicro/payment-srv/handler"
	"github.com/songxuexian/gogomicro/payment-srv/subscriber"

	payment "github.com/songxuexian/gogomicro/payment-srv/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.payment", service.Server(), new(subscriber.Payment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
