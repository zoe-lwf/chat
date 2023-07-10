package main

import (
	"chat/config"
	"chat/pkg/db"
	"chat/router"
)

func main() {
	//配置初始化
	config.InitConfig("./app.yaml")
	//初始化mysql
	db.InitMysql(config.GlobalConfig.MySQL.DNS)
	//初始化redis
	db.InitRedis(config.GlobalConfig.Redis.Addr, config.GlobalConfig.Redis.Password)
	// router启动
	router.HTTPRouter()
}
