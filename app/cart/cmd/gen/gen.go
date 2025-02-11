package main

import (
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/model"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/mysql"
	"gorm.io/gen"
)

func main() {

	mysql.Init()
	db := mysql.DB

	g := gen.NewGenerator(gen.Config{
		OutPath: "biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)
	g.ApplyInterface(func(model.Querier) {}, model.CartItem{})
	g.Execute()
}
