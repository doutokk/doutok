package dal

import (
	"github.com/doutokk/doutok/app/product/biz/dal/mysql"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
