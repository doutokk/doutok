package main

import (
	"fmt"
	"github.com/doutokk/doutok/app/auth/infra/casbin"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/auth/biz/dal/mysql"
	"github.com/doutokk/doutok/app/auth/conf"
	mysqldb "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var c = conf.GetConf()

func main() {
	// connect to mysql manually to check and create database
	dsn := "%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysqldb.Open(fmt.Sprintf(dsn, c.MySQL.Username, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	var count int
	dbName := conf.GetConf().Kitex.Service
	db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count)
	if count == 0 {
		klog.Warn("Database not found, creating database")
		db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", dbName))
	}

	// migrate the database
	mysql.Init()

	err = mysql.DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&gormadapter.CasbinRule{})
	casbin.InitPolicy()

	if err != nil {
		panic(err)
	}
}
