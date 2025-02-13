package main

import (
	"github.com/doutokk/doutok/app/cart/biz/dal/model"
	"github.com/doutokk/doutok/app/cart/biz/dal/mysql"
)

func main() {
	mysql.Init()
	err := mysql.DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		panic(err)
	}
}
