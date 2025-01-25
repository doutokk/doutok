package main

import (
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/model"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/mysql"
)

func main() {
	mysql.Init()
	err := mysql.DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		panic(err)
	}
}
