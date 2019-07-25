package orders

import (
	"fmt"
	"github.com/micro/go-micro/client"
	proto "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
	"sync"
)

var (
	s         *service
	invClinet invS.InventoryService
	m         *sync.RWMutex
)

type service struct {
}

type Service interface {
	New(bookId, userId int64) (orderId int64, err error)
	GetOrder(orderId int64) (order *proto.Order, err error)
	UpdateOrderState(orderId int64, state int) (err error)
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
	invClinet = invS.NewInventoryService("sxx.micro.book.srv.inventory", client.DefaultClient)
	s = &service{}
}
