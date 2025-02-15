package dal

import (
	"github.com/doutokk/doutok/app/bff/biz/dal/mysql"
	"github.com/doutokk/doutok/app/bff/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
