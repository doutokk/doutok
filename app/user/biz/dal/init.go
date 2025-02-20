package dal

import (
	"github.com/doutokk/doutok/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
