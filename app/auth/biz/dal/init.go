package dal

import (
	"github.com/doutokk/doutok/app/auth/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
