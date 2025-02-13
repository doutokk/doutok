package mysql

import (
	"fmt"
	"github.com/doutokk/doutok/app/payment/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
	c   = conf.GetConf()
)

func Init() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(dsn, c.MySQL.Username, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.Kitex.Service)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}
