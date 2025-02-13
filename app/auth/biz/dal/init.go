package dal

import (
	"github.com/doutokk/doutok/app/auth/biz/dal/mysql"
	"github.com/doutokk/doutok/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
