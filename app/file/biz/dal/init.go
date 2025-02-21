package dal

import (
	"github.com/doutokk/doutok/app/file/biz/dal/mysql"
	"github.com/doutokk/doutok/app/file/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
