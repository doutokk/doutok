package dal

import (
	"github.com/doutokk/doutok/app/gateway/biz/dal/mysql"
	"github.com/doutokk/doutok/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
