package main

import (
	"chat/config"
	"chat/pkg/db"
	"chat/router"
)

func main() {
	//配置初始化
	config.InitConfig("./app.yaml")
	//初始化
	db.InitMysql(config.GlobalConfig.MySQL.DNS)

	// router启动
	router.HTTPRouter()
}
