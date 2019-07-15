package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/config"
)

func initMysql() {
	var err error
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatalf("[Init mysql]" + err.Error())
		panic(err)
	}

	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatalf("[Init mysql]" + err.Error())
		panic(err)
	}
}
