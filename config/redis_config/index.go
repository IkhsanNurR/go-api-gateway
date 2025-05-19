package redisconfig

import (
	appconfig "api-gateway/config/app_config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	rdb *redis.Client
)

// InitRedis initializes Redis connection
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: appconfig.RedisHost,
		// Password: appconfig.RedisPassword, // No password set
		DB: 0, // Use default DB
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	fmt.Println("✅ Redis connected")
}

// Set sets a key with a value and expiration
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(Ctx, key, value, expiration).Err()
}

// Get retrieves the value of a key
func Get(key string) (string, error) {
	return rdb.Get(Ctx, key).Result()
}

// Delete removes a key from Redis
func Delete(key string) error {
	return rdb.Del(Ctx, key).Err()
}
