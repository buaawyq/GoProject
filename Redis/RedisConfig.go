package Redis

import (
	"github.com/go-redis/redis"
	"time"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        "192.168.1.103:6379", // Redis地址
		Password:    "",                   // Redis账号
		DB:          0,                    // Redis库
		PoolSize:    10,                   // Redis连接池大小
		MaxRetries:  3,                    // 最大重试次数
		IdleTimeout: 10 * time.Second,     // 空闲链接超时时间
	})
	return client
}
