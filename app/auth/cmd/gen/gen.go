package main

import (
	"github.com/doutokk/doutok/app/auth/biz/dal/mysql"
	"gorm.io/gen"
)

func main() {
	// connect to database
	mysql.Init()
	db := mysql.DB

	// create the code generator
	g := gen.NewGenerator(gen.Config{
		OutPath: "biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// bind the database to code generator
	g.UseDB(db)

	g.ApplyInterface(func() {})
	g.Execute()
}
