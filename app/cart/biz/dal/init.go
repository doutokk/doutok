package dal

import (
	"github.com/doutokk/doutok/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
