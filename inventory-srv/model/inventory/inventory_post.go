package inventory

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/common"
	"github.com/songxuexian/gogomicro/basic/db"
	proto "github.com/songxuexian/gogomicro/inventory-srv/proto/inventory"
)

func (s *service) Sell(bookId, userId int64) (id int64, err error) {
	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Logf("[Shell] Begin transaction fail", err.Error())
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	querySQL := `SELECT id, book_id, unit_price, stock, version FROM inventory WHERE book_id = ?`

	inv := &proto.Inv{}
	updateSQL := `UPDATE inventory SET stock=?, version=?  WHERE book_id = ? AND version=?`
	var defuctInv func() error
	defuctInv = func() (errIn error) {
		errIn = tx.QueryRow(querySQL, bookId).Scan(&inv.Id, &inv.BookId, &inv.UnitPrice, &inv.Stock, &inv.Version)
		if errIn != nil {
			log.Logf("[Shell]select date fail,err: %s", errIn.Error())
			return
		}
		if inv.Stock < 1 {
			errIn = fmt.Errorf("[Shell] inventory not enought")
			log.Logf("[Shell]select date fail,err: %s", errIn.Error())
			return
		}
		result, errIn := tx.Exec(updateSQL, inv.Stock-1, inv.Version+1, bookId, inv.Version)
		if errIn != nil {
			log.Logf("[Shell]update inventory fail,err: %s", errIn.Error())
			return
		}
		if affected, _ := result.RowsAffected(); affected == 0 {
			log.Logf("[Shell]update inventory fail,version %d out date,will be try again", inv.Version)
			_ = defuctInv()
		}
		return
	}
	err = defuctInv()
	if err != nil {
		log.Logf("[Shell] Sub inventory fail,err: %s", err.Error())
		return
	}
	insertSQL := `INSERT inventory_history (book_id, user_id, state) VALUE (?, ?, ?) `
	r, err := tx.Exec(insertSQL, bookId, userId, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[Sell] 新增销存记录失败，err：%s", err)
		return
	}

	// 返回历史记录id，作为流水号使用
	id, _ = r.LastInsertId()

	// 忽略error
	tx.Commit()
	return
}

// Confirm 确认销存
func (s *service) Confirm(id int64, state int) (err error) {
	updateSQL := `UPDATE inventory_history SET state = ? WHERE id = ?;`

	// 获取数据库
	o := db.GetDB()

	// 更新
	_, err = o.Exec(updateSQL, state, id)
	if err != nil {
		log.Logf("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
