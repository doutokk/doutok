package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
