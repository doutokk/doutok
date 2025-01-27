package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
