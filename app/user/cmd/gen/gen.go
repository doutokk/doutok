package main

import (
	"github.com/doutokk/doutok/app/user/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:2048711712P!@tcp(127.0.0.1:3306)/doutok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func main() {
	// connect to database
	//err := godotenv.Load(".env")
	//if err != nil {
	//	dir, _ := os.Getwd()
	//	fmt.Println("Current working directory:", dir)
	//	panic(".env file not found")
	//	return
	//}
	//mysql.Init()
	//db := mysql.DB
	db := initDB()

	// create the code generator
	g := gen.NewGenerator(gen.Config{
		OutPath: "biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// bind the database to code generator
	g.UseDB(db)

	g.ApplyInterface(func(model.Querier) {}, &model.User{})
	g.Execute()
}
