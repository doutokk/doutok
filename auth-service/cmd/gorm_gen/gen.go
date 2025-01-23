package main

import (
	"douyin-commerce/auth-service/app/model"
	"douyin-commerce/auth-service/pkg/mysql"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: "auth_service/biz/dal/mysql/gen", Mode: gen.WithDefaultQuery | gen.WithQueryInterface})

	g.UseDB(mysql.InitDB())

	g.ApplyBasic(model.Token{})

	g.Execute()
}
