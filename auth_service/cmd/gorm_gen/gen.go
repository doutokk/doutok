package main

import (
	"douyin-commerce/auth_service/biz/dal/mysql"
	"douyin-commerce/auth_service/biz/dal/mysql/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: "auth_service/biz/dal/mysql/gen", Mode: gen.WithDefaultQuery | gen.WithQueryInterface})

	g.UseDB(mysql.InitDB())

	g.ApplyBasic(model.Token{})

	g.Execute()
}
