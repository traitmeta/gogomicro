package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/orders-srv/model/orders"
)

var (
	ordersService orders.Service
)

func Init() {
	ordersService, _ = orders.GetService()
}

func PayOrder(ctx context.Context, event *payS.PayEvent) (err error) {
	log.Logf("[PayOrder] Received pay message: %d, %d", event.OrderId, event.State)
	err = ordersService.UpdateOrderState(event.OrderId, int(event.State))
	if err != nil {
		log.Logf("[PayOrder] Received pay message,update state fail, %s", err)
		return
	}
	return
}
