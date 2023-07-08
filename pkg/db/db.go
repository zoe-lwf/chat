package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DB *gorm.DB
)

func InitMysql(dataSource string) {
	//TODO logger

	//连接
	db, err := gorm.Open(mysql.Open(dataSource),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	////数据库DB
	sqlDb, err := db.DB()

	if err != nil {
		panic(err)
	}
	//用于设置连接池中空闲连接的最大数量
	sqlDb.SetMaxOpenConns(20)
	//设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(30)
	//设置了连接可复用的最大时间
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 赋值
	DB = db
	////logger.Logger.Info("mysql init ok")
	fmt.Println("mysql init ok")
}
