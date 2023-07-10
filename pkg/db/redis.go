package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis(addr, password string) {
	//logger.Logger.Debug("Redis init ...")
	RDB = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password, // no password set
		DB:           0,        // use default DB
		PoolSize:     30,
		MinIdleConns: 30,
	})
	err := RDB.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis init ok")
	//logger.Logger.Debug("Redis init ok")
}

// Publish 发送消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := RDB.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RDB.PSubscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
