package dal

import (
	"github.com/PengJingzhao/douyin-commerce/app/user/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
