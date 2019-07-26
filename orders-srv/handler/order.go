package handler

import (
	"context"
	"github.com/songxuexian/gogomicro/orders-srv/model/orders"

	"github.com/micro/go-micro/util/log"

	proto "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
)

var (
	ordersService orders.Service
)

type Orders struct{}

func Init() {
	ordersService, _ = orders.GetService()
}

func (e *Orders) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	orderId, err := ordersService.New(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}
	rsp.Order = &proto.Order{
		Id: orderId,
	}

	return
}

func (e *Orders) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[GetOrder] Received get order request, %d", req.OrderId)
	rsp.Order, err = ordersService.GetOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}
	rsp.Success = true
	return
}
