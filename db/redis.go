package db

import (
	"fmt"

	"gopkg.in/redis.v5"
)

// Redis - Redis client
var Redis *redis.Client

// MakeRedisConn - Create redis connection
func MakeRedisConn(url string, pwd string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pwd,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println("Redis started ? ping...", pong, err)
	return client
}
