package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/auth/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
