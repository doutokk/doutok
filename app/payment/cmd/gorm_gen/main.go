package main

import (
	"github.com/PengJingzhao/douyin-commerce/app/payment/biz/dal/model"

	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.ApplyInterface(func(model.Querier) {}, model.User{})
	g.Execute()
}
