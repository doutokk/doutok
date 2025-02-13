package dal

import (
	"github.com/doutokk/doutok/app/checkout/biz/dal/mysql"
	"github.com/doutokk/doutok/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
