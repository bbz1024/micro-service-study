package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: redis
 * @Date:
 * @Desc: ...
 *
 */

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "192.168.40.129:6379",
		DB:   8,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	// 初始化用户余额
	RedisClient.HSet(context.Background(), fmt.Sprintf(UserKey, 1), "balance", 100)

}
