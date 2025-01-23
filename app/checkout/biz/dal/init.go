package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/checkout/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
