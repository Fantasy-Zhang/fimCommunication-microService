package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr, //"localhost:6379",
		Password: pwd,  //"redis", // no password set
		DB:       db,   // use default DB
		PoolSize: 100,  //
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("redis连接失败")
		return
	}
	return rdb
}
