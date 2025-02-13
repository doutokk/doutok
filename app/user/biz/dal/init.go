package dal

import (
	"github.com/doutokk/doutok/app/user/biz/dal/mysql"
	"github.com/doutokk/doutok/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
