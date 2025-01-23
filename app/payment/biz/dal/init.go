package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/payment/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
