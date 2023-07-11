package db

import (
	"chat/config"
	"context"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	config.InitConfig("./app.yaml")
	//初始化redis
	InitRedis(config.GlobalConfig.Redis.Addr, config.GlobalConfig.Redis.Password)

	msg, err := RDB.Set(context.Background(), "712", "cao", 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)

}
