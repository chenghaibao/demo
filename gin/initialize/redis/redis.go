package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hb_gin/config"
	"time"
)

var Redis *redis.Client

func NewRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr, // use default Addr
		Password: "",                     // no password set
		DB:       0,                      // use default DB
		PoolSize: 100, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Redis.Ping(ctx).Result()
}
