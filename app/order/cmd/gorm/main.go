package main

import (
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/model"
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/mysql"
)

func main() {
	mysql.Init()
	err := mysql.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
