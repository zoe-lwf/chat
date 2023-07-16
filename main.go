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
	//初始化MQ
	// mq.InitMessageMQ(config.GlobalConfig.RabbitMQ.URL)
	// router启动 http服务
	go router.HTTPRouter()

	// 启动 websocket 服务
	router.WSRouter()
}
