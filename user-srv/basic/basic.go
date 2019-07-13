package basic

import (
	"github.com/songxuexian/gogomicro/user-srv/basic/config"
	"github.com/songxuexian/gogomicro/user-srv/basic/db"
)

func Init(){
	config.Init()
	db.Init()
}