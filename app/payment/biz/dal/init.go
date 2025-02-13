package dal

import (
	"github.com/doutokk/doutok/app/payment/biz/dal/mysql"
	"github.com/doutokk/doutok/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
