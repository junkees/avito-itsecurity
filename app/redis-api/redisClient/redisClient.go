package redisClient

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func GetConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("PASSWORD"),
		DB:       0,
	})
}
