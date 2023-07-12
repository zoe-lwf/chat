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

const (
	PublicKey = "chat_user"
)

// Publish 发送消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := RDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RDB.Subscribe(ctx, channel)
	//ch := sub.Channel()
	//for msg := range ch {
	//	fmt.Println(msg.Channel, msg.Payload)
	//}
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return msg.Payload, err
}
