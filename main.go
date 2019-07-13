package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
	proto "github.com/songxuexian/gogomicro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}


func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)
	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
