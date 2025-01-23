package dal

import (
	"github.com/PengJingzhao/douyin-commerce/cart-service/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/cart-service/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
