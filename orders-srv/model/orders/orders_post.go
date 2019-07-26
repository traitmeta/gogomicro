package orders

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/common"
	"github.com/songxuexian/gogomicro/basic/db"
)

func (s *service) New(bookId, userId int64) (orderId int64, err error) {
	resp, err := invClinet.Sell(context.TODO(), &invS.Request{
		BookID: bookId, UserId: userId,
	})
	if err != nil {
		log.Logf("[New] Sell call inventory service fail: %s", err.Error())
	}

	o := db.GetDB()
	insertSQL := `INSERT orders (user_id,book_id,inv_his_id,state) VALUE (?,?,?,?)`

	r, err := o.Exec(insertSQL, userId, bookId, resp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[New] add order fail, err: %s", err.Error())
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSLQ := `UPDATE orders SET state = ? WHERE id =?`
	o := db.GetDB()
	_, err = o.Exec(updateSLQ, state, orderId)
	if err != nil {
		log.Logf("[Update] update fail, err: %s", err.Error())
	}
	return
}
