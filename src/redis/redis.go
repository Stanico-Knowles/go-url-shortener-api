package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func ConnectRedis() {
	ctx := context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failure: %v", err)
	} else {
		log.Println("Redis connection successful")
	}
}

func SetItem(key string, value string) error {
	ctx := context.Background()
	return redisClient.Set(ctx, key, value, 0).Err()
}

func GetItem(key string) (string, error) {
	ctx := context.Background()
	return redisClient.Get(ctx, key).Result()
}

func DeleteItem(key string) {
	ctx := context.Background()
	redisClient.Del(ctx, key)
}

func CloseRedis() {
	if redisClient != nil {
		redisClient.Close()
	}
}
