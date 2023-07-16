package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
	log.Println(redisClient.Ping(ctx))
	return redisClient
}