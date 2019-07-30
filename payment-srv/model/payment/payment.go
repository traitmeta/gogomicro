package payment

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/songxuexian/gogomicro/basic/common"
	invS "github.com/songxuexian/gogomicro/inventory-srv/proto/inventory"
	ordS "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
	"sync"
)

var (
	s            *service
	m            sync.RWMutex
	ordSClient   ordS.OrdersService
	payPublisher micro.Publisher
	invClient    invS.InventoryService
)

type service struct {
}

type Service interface {
	PayOrder(orderId int64) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService not init")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = invS.NewInventoryService("sxx.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("sxx.micro.book.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
