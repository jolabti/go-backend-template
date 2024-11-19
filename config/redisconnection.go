package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/firmanJS/fiber-with-mongo/config"
	"github.com/go-redis/redis/v8"
)

// RedisWrapper struct holds the Redis client and context
type RedisWrapper struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisWrapper creates a new Redis client and context
func NewRedisWrapper(db int) *RedisWrapper {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	// Test the connection
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisWrapper{
		client: client,
		ctx:    ctx,
	}
}

// Set writes a key-value pair to Redis
func (rw *RedisWrapper) Set(key string, value string, expiration time.Duration) error {
	err := rw.client.Set(rw.ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not set key %s: %v", key, err)
	}
	return nil
}

func (rw *RedisWrapper) Delete(key string) error {
	err := rw.client.Del(rw.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("could not delete key %s: %v", key, err)
	}
	return nil
}

// Get reads the value for a given key from Redis
func (rw *RedisWrapper) Get(key string) (string, error) {
	val, err := rw.client.Get(rw.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		}
		return "", fmt.Errorf("could not get key %s: %v", key, err)
	}
	return val, nil
}
