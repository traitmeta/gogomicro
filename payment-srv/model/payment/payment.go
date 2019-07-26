package payment

import (
	"github.com/micro/go-micro"
	ordS "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
	"sync"
)

var (
	s            *service
	m            sync.RWMutex
	ordSClient   ordS.OrdersService
	payPublisher micro.Publisher
)

type service struct {
}

type Service interface {
	PayOrder(orderId int64) (err error)
}

func Init() {

}
