package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
