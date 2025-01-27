package mysql

import (
	"fmt"
	"github.com/PengJingzhao/douyin-commerce/app/auth/biz/model"
	"github.com/PengJingzhao/douyin-commerce/app/user/conf"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// AutoMigrate User Table
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
}
