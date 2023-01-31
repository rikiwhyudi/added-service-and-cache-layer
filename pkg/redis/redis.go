package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func RedisClientInit() *redis.Client {

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println("Redis connected...")

	return RDB

}
