package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client *redis.Client

func InitRedisClient() {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Panicf("Could not connect to redis client\nError: %v", err)
	}

	client = redis.NewClient(opts)
}

func Set(key, value string) error {
	hourToDay := time.Hour * 24
	err := client.Set(ctx, key, value, hourToDay*14).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
