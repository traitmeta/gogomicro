package basic

import (
	"github.com/songxuexian/gogomicro/basic/config"
	"github.com/songxuexian/gogomicro/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
