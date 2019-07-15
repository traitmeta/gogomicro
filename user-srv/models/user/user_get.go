package user

import (
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/db"
	proto "github.com/songxuexian/gogomicro/user-srv/proto/user"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	mysqlDB := db.GetDB()
	if mysqlDB == nil {
		log.Log("[QueryUserByName] Get DB fail")
		return
	}

	ret = &proto.User{}

	// 查询
	queryRow := mysqlDB.QueryRow(queryString, userName)
	if queryRow == nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	err = queryRow.Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
