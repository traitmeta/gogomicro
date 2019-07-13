package db

import (
	"database/sql"
	"github.com/go-log/log"
	"github.com/songxuexian/gogomicro/user-srv/basic/config"
)

func initMysql() {
	var err error
	mysqlDB, err := sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Logf(err.Error())
		panic(err)
	}
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	if err := mysqlDB.Ping(); err != nil {
		log.Logf(err.Error())
		panic(err)
	}
}
