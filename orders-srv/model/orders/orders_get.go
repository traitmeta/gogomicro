package orders

import (
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/db"
	proto "github.com/songxuexian/gogomicro/orders-srv/proto/orders"
)

func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	o := db.GetDB()

	err = o.QueryRow("SELECT id, user_id, book_id,inv_his_id, state FROM orders WHERE id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Logf("[GetOrder] get data fail, err: %s", err)
		return
	}
	return
}
