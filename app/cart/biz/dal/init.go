package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
