package dal

import (
	"github.com/doutokk/doutok/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
