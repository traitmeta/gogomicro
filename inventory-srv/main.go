package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/songxuexian/gogomicro/inventory-srv/handler"
	"github.com/songxuexian/gogomicro/inventory-srv/subscriber"

	inventory "github.com/songxuexian/gogomicro/inventory-srv/proto/inventory"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("sxx.micro.book.srv.inventory"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.inventory", service.Server(), new(subscriber.Inventory))

	// Register Function as Subscriber
	micro.RegisterSubscriber("sxx.micro.book.srv.inventory", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
