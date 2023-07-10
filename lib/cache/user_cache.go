package cache

import (
	"chat/pkg/db"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	userOnlinePrefix = "user_online_" // 用户在线状态设置
	ttl1D            = 24 * 60 * 60   //s  1天
)

func getUserKey(userId uint64) string {
	return fmt.Sprintf("%s%d", userOnlinePrefix, ttl1D)
}

// SetUserOnline 设置用户在线
func SetUserOnline(userid uint64, addr string) error {
	key := getUserKey(userid)
	_, err := db.RDB.Set(context.Background(), key, addr, ttl1D*time.Second).Result()
	if err != nil {
		fmt.Println("[设置用户在线] 错误, err:", err)
		return err
	}
	return nil
}

// GetUserOnline 获取用户在线地址
func GetUserOnline(userid uint64) (string, error) {
	key := getUserKey(userid)
	addr, err := db.RDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("[获取用户在线] 错误，err:", err)
		return "", err
	}
	return addr, nil
}

// DelUserOnline 删除用户在线信息（存在即在线）
func DelUserOnline(userid uint64) error {
	key := getUserKey(userid)
	_, err := db.RDB.Del(context.Background(), key).Result()
	if err != nil {
		fmt.Println("[删除用户在线] 错误, err:", err)
		return err
	}
	return nil
}
