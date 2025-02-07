package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
