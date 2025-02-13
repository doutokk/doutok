package dal

import (
	"github.com/doutokk/doutok/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
