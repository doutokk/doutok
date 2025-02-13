package dal

import (
	"github.com/doutokk/doutok/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
